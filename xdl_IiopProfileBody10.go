// Code generated by me. DO NOT EDIT.

package goIdlIiop

import __goIdlIop__ "github.com/bhbosman/goIdlIop"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"

// Struct declaration: "IIOP::ProfileBody_1_0", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type IiopProfileBody10 struct {
	__goidl__.IdlObject
	IiopVersion IiopVersion `json:"IiopVersion"`
	Host string `json:"Host"`
	Port uint16 `json:"Port"`
	ObjectKey __goIdlIop__.IopObjectKey `json:"ObjectKey"`
}

//noinspection GoSnakeCaseUsage
const IiopProfileBody10Id_Const = "IDL:omg.org/IIOP/ProfileBody_1_0:1.0"

func (self *IiopProfileBody10) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IiopProfileBody10) GoString() string {
	return self.String()
}

func (self *IiopProfileBody10) ReadValue(stream __goidl__.IReadAny) error {
	var err error
	err = self.IdlObject.ReadValue(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.IiopVersion.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(StringType)
	self.Host, err = __goidl__.IdlStringHelper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(UnsignedShortType)
	self.Port, err = __goidl__.IdlUInt16Helper.Read(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.ObjectKey.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopProfileBody10) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopProfileBody10) Write(stream __goidl__.IWriteAny) error {
	var err error
	err = self.IdlObject.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.IiopVersion.Write(stream)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(StringType)
	err = __goidl__.IdlStringHelper.Write(stream, self.Host)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(UnsignedShortType)
	err = __goidl__.IdlUInt16Helper.Write(stream, self.Port)
	if err != nil {
		return err
	}
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.ObjectKey.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IiopProfileBody10_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IiopProfileBody10Helper = IiopProfileBody10_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IiopProfileBody10Id_Const,
			__goidl__.StructType,
			"IIOP.idl",
			"xdl_IiopProfileBody10.go",
			func() __goidl__.IIdlObject {
				return &IiopProfileBody10{}
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return &IiopProfileBody10{}
			},
			__reflect__.TypeOf((*IiopProfileBody10)(nil))))
}
