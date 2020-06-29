// Code generated by me. DO NOT EDIT.

package goIdlIiop

import __goIdlCorba__ "github.com/bhbosman/goIdlCorba"
import __goIdlIop__ "github.com/bhbosman/goIdlIop"
import __goidl__ "github.com/bhbosman/goidl"
import __json__ "encoding/json"
import __reflect__ "reflect"
import __yaccidl__ "github.com/bhbosman/yaccidl"

// Struct declaration: "IIOP::ProfileBody_1_1", generatedBy by: "WriteStructDcl"
// Exception Decl: false
type IiopProfileBody11 struct {
	__goidl__.IdlObject
	IiopVersion IiopVersion `json:"IiopVersion"`
	Host string `json:"Host"`
	Port uint16 `json:"Port"`
	ObjectKey __goIdlIop__.IopObjectKey `json:"ObjectKey"`
	Components __goIdlIop__.IopTaggedComponentSeq `json:"Components"`
}

//noinspection GoSnakeCaseUsage
const IiopProfileBody11Id_Const = "IDL:omg.org/IIOP/ProfileBody_1_1:1.0"

//noinspection GoUnusedExportedFunction
func NewDefaultIiopProfileBody11() *IiopProfileBody11 {
	//noinspection GoRedundantConversion
	return &IiopProfileBody11{
		IdlObject: __goidl__.IdlObject{
			Id: IiopProfileBody11Id_Const,
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "iiop_version", Type: "IIOP::Version" Scope: "IdlStruct"
		IiopVersion: IiopVersion{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "major", Type: "IdlOctet" Scope: "IdlOctetKind"
		Major: 0,
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "minor", Type: "IdlOctet" Scope: "IdlOctetKind"
		Minor: 0,
	},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "host", Type: "IdlString" Scope: "StringType"
		Host: "",
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "port", Type: "IdlUnsignedShort" Scope: "UnsignedShortType"
		Port: 0,
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "object_key", Type: "IOP::ObjectKey" Scope: "IdlStruct"
		ObjectKey: __goIdlIop__.IopObjectKey{
	CorbaOctetSeq: __goIdlCorba__.CorbaOctetSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IdlOctet" Scope: "IdlOctetKind"
		Array: 0,
	},
	},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "components", Type: "IOP::TaggedComponentSeq" Scope: "IdlStruct"
		Components: __goIdlIop__.IopTaggedComponentSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IOP::TaggedComponent" Scope: "IdlStruct"
		Array: __goIdlIop__.IopTaggedComponent{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "tag", Type: "IOP::ComponentId" Scope: "IdlTypedef"
		Tag: 0,
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "component_data", Type: "IOP::ComponentData" Scope: "IdlStruct"
		ComponentData: __goIdlIop__.IopComponentData{
	CorbaOctetSeq: __goIdlCorba__.CorbaOctetSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IdlOctet" Scope: "IdlOctetKind"
		Array: 0,
	},
	},
	},
	},
	}
}

//noinspection GoUnusedExportedFunction
func NewIiopProfileBody11(
	//"StructDclAllParamsConstructorParamService"
	IiopVersion IiopVersion,
	//"StructDclAllParamsConstructorParamService"
	Host string,
	//"StructDclAllParamsConstructorParamService"
	Port uint16,
	//"StructDclAllParamsConstructorParamService"
	ObjectKey __goIdlIop__.IopObjectKey,
	//"StructDclAllParamsConstructorParamService"
	Components __goIdlIop__.IopTaggedComponentSeq) *IiopProfileBody11 {
	//noinspection GoRedundantConversion
	return &IiopProfileBody11{
		IdlObject: __goidl__.IdlObject{
			Id: IiopProfileBody11Id_Const,
		},
		IiopVersion: IiopVersion,
		Host: Host,
		Port: Port,
		ObjectKey: ObjectKey,
		Components: Components,
	}
}

//noinspection GoUnusedExportedFunction
func NewRandomIiopProfileBody11(randomGenerator __goidl__.IRandomDataGenerator) *IiopProfileBody11 {
	//noinspection GoRedundantConversion
	return &IiopProfileBody11{
		IdlObject: __goidl__.IdlObject{
			Id: IiopProfileBody11Id_Const,
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "iiop_version", Type: "IIOP::Version" Scope: "IdlStruct"
		IiopVersion: IiopVersion{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "major", Type: "IdlOctet" Scope: "IdlOctetKind"
		Major: randomGenerator.OctetTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Major"),
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "minor", Type: "IdlOctet" Scope: "IdlOctetKind"
		Minor: randomGenerator.OctetTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Minor"),
	},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "host", Type: "IdlString" Scope: "StringType"
		Host: randomGenerator.StringTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Host"),
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "port", Type: "IdlUnsignedShort" Scope: "UnsignedShortType"
		Port: randomGenerator.UnsignedShortTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Port"),
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "object_key", Type: "IOP::ObjectKey" Scope: "IdlStruct"
		ObjectKey: __goIdlIop__.IopObjectKey{
	CorbaOctetSeq: __goIdlCorba__.CorbaOctetSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IdlOctet" Scope: "IdlOctetKind"
		Array: randomGenerator.OctetTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Array"),
	},
	},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "components", Type: "IOP::TaggedComponentSeq" Scope: "IdlStruct"
		Components: __goIdlIop__.IopTaggedComponentSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IOP::TaggedComponent" Scope: "IdlStruct"
		Array: __goIdlIop__.IopTaggedComponent{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "tag", Type: "IOP::ComponentId" Scope: "IdlTypedef"
		Tag: randomGenerator.UnsignedLongTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Tag"),
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "component_data", Type: "IOP::ComponentData" Scope: "IdlStruct"
		ComponentData: __goIdlIop__.IopComponentData{
	CorbaOctetSeq: __goIdlCorba__.CorbaOctetSeq{
		IdlObject: __goidl__.IdlObject{
			Id: "ddd",
		},
		// "StructDclDefaultConstructorMemberDefaultValueService", MemberName: "Array", Type: "IdlOctet" Scope: "IdlOctetKind"
		Array: randomGenerator.OctetTypeValue(__reflect__.TypeOf((*IiopProfileBody11)(nil)), "Array"),
	},
	},
	},
	},
	}
}

func (self *IiopProfileBody11) String() string {
	b, err := __json__.Marshal(self)
	if err != nil {
		return ""
	}
	return string(b)
}

func (self *IiopProfileBody11) GoString() string {
	return self.String()
}

func (self *IiopProfileBody11) ReadValue(stream __goidl__.IReadAny) error {
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
	// WriteStructHelper::WriteStructMemberExtractValue(IdlStruct)
	err = self.Components.Read(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopProfileBody11) Read(stream __goidl__.IReadAny) error {
	err := self.ReadValue(stream)
	if err != nil {
		return err
	}
	return nil
}

func (self *IiopProfileBody11) Write(stream __goidl__.IWriteAny) error {
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
	// WriteStructHelper::WriteStructMemberInsert(IdlStruct)
	err = self.Components.Write(stream)
	if err != nil {
		return err
	}
	return nil
}

//noinspection GoSnakeCaseUsage
type IiopProfileBody11_Helper struct {
}


//noinspection GoUnusedGlobalVariable
var IiopProfileBody11Helper = IiopProfileBody11_Helper{}

func init() {
	__goidl__.AddRegistration(
		__goidl__.NewRegistrationInformation(
			IiopProfileBody11Id_Const,
			__yaccidl__.IdlStruct,
			"IIOP.idl.ProfileBody_1_1.go",
			func() __goidl__.IIdlObject {
				return NewDefaultIiopProfileBody11()
			},
			func(generator __goidl__.IRandomDataGenerator) __goidl__.IIdlObject {
				return NewRandomIiopProfileBody11(generator)
			},
			__reflect__.TypeOf((*IiopProfileBody11)(nil))))
}
