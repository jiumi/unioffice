//
// Copyright 2020 FoxyUtils ehf. All rights reserved.
//
// This is a commercial product and requires a license to operate.
// A trial license can be obtained at https://unidoc.io
//
// DO NOT EDIT: generated by unitwist Go source code obfuscator.
//
// Use of this source code is governed by the UniDoc End User License Agreement
// terms that can be accessed at https://unidoc.io/eula/

// Package memstore implements tempStorage interface
// by using memory as a storage
package memstore ;import (_ga "encoding/hex";_dg "errors";_eb "fmt";_c "github.com/unidoc/unioffice/common/tempstorage";_e "io";_d "io/ioutil";_ec "math/rand";_g "sync";);

// ReadAt reads from the underlying memDataCell at an offset provided in order to implement ReaderAt interface.
// It does not affect f.readOffset.
func (_a *memFile )ReadAt (p []byte ,readOffset int64 )(int ,error ){_ag :=_a ._cb ._fab ;_db :=int64 (len (p ));if _db > _ag {_db =_ag ;p =p [:_db ];};if readOffset >=_ag {return 0,_e .EOF ;};_fa :=readOffset +_db ;if _fa >=_ag {_fa =_ag ;};_ff :=copy (p ,_a ._cb ._gg [readOffset :_fa ]);return _ff ,nil ;};type memFile struct{_cb *memDataCell ;_b int64 ;};

// Open returns tempstorage File object by name
func (_ebg *memStorage )Open (path string )(_c .File ,error ){_gc ,_ae :=_ebg ._fb .Load (path );if !_ae {return nil ,_dg .New (_eb .Sprintf ("\u0043\u0061\u006eno\u0074\u0020\u006f\u0070\u0065\u006e\u0020\u0074\u0068\u0065\u0020\u0066\u0069\u006c\u0065\u0020\u0025\u0073",path ));};return &memFile {_cb :_gc .(*memDataCell )},nil ;};

// Close is not applicable in this implementation
func (_da *memFile )Close ()error {return nil };type memDataCell struct{_ddf string ;_gg []byte ;_fab int64 ;};func _cga (_ddad int )(string ,error ){_gaa :=make ([]byte ,_ddad );if _ ,_bce :=_ec .Read (_gaa );_bce !=nil {return "",_bce ;};return _ga .EncodeToString (_gaa ),nil ;};

// RemoveAll removes all files according to the dir argument prefix
func (_fc *memStorage )RemoveAll (dir string )error {_fc ._fb .Range (func (_cd ,_gd interface{})bool {_fc ._fb .Delete (_cd );return true });return nil ;};

// TempDir creates a name for a new temp directory using a pattern argument
func (_bd *memStorage )TempDir (pattern string )(string ,error ){return _cg (pattern ),nil };

// TempFile creates a new empty file in the storage and returns it
func (_aeb *memStorage )TempFile (dir ,pattern string )(_c .File ,error ){_dda :=dir +"\u002f"+_cg (pattern );_aad :=&memDataCell {_ddf :_dda ,_gg :[]byte {}};_bb :=&memFile {_cb :_aad };_aeb ._fb .Store (_dda ,_aad );return _bb ,nil ;};func _cg (_fcd string )string {_ab ,_ :=_cga (6);return _fcd +_ab };

// Name returns the filename of the underlying memDataCell
func (_fg *memFile )Name ()string {return _fg ._cb ._ddf };

// Add reads a file from a disk and adds it to the storage
func (_af *memStorage )Add (path string )error {_ ,_bef :=_af ._fb .Load (path );if _bef {return nil ;};_ge ,_ad :=_d .ReadFile (path );if _ad !=nil {return _ad ;};_af ._fb .Store (path ,&memDataCell {_ddf :path ,_gg :_ge ,_fab :int64 (len (_ge ))});return nil ;};type memStorage struct{_fb _g .Map };

// Read reads from the underlying memDataCell in order to implement Reader interface
func (_fd *memFile )Read (p []byte )(int ,error ){_bc :=_fd ._b ;_be :=_fd ._cb ._fab ;_bg :=int64 (len (p ));if _bg > _be {_bg =_be ;p =p [:_bg ];};if _bc >=_be {return 0,_e .EOF ;};_ecd :=_bc +_bg ;if _ecd >=_be {_ecd =_be ;};_dd :=copy (p ,_fd ._cb ._gg [_bc :_ecd ]);_fd ._b =_ecd ;return _dd ,nil ;};

// Write writes to the end of the underlying memDataCell in order to implement Writer interface
func (_aa *memFile )Write (p []byte )(int ,error ){_aa ._cb ._gg =append (_aa ._cb ._gg ,p ...);_aa ._cb ._fab +=int64 (len (p ));return len (p ),nil ;};

// SetAsStorage sets temp storage as a memory storage
func SetAsStorage (){_ac :=memStorage {_fb :_g .Map {}};_c .SetAsStorage (&_ac )};