// Code generated by me. DO NOT EDIT.

package goIdlIiop

import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "IIOP::Version", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type IiopVersion struct {
	__goidl__.IdlObject
	Major byte `json:"Major"`
	Minor byte `json:"Minor"`
}

//noinspection GoSnakeCaseUsage
const IiopVersionId_Const = "IDL:omg.org/IIOP/Version:1.0"

func (self *IiopVersion) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IiopVersion) GoString() string {
	return self.String()
}

func (self *IiopVersion) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlOctetKind)
	self.Major, err = __goidl__.IdlOctetHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlOctetKind)
	self.Minor, err = __goidl__.IdlOctetHelper.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopVersion) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopVersion) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlOctetKind)
	err = __goidl__.IdlOctetHelper.Write(stream, self.Major)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlOctetKind)
	err = __goidl__.IdlOctetHelper.Write(stream, self.Minor)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IiopVersion_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IiopVersionHelper = IiopVersion_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IiopVersionId_Const,
			__goidl__.StructType,
			"IIOP.idl",
			"xdl_IiopVersion.go",
			func() __goidl__.IIdlObject {
				return &IiopVersion{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IiopVersion{}
			},
			__reflect__.TypeOf((*IiopVersion)(nil))))
}
