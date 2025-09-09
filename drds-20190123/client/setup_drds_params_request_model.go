// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetupDrdsParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*SetupDrdsParamsRequestData) *SetupDrdsParamsRequest
	GetData() []*SetupDrdsParamsRequestData
	SetDrdsInstanceId(v string) *SetupDrdsParamsRequest
	GetDrdsInstanceId() *string
	SetParamLevel(v string) *SetupDrdsParamsRequest
	GetParamLevel() *string
	SetRegionId(v string) *SetupDrdsParamsRequest
	GetRegionId() *string
}

type SetupDrdsParamsRequest struct {
	// This parameter is required.
	Data []*SetupDrdsParamsRequestData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The ID of the PolarDB-X 1.0 instance for which you want to configure parameters.
	//
	// This parameter is required.
	//
	// example:
	//
	// drdsjiii1b49****
	DrdsInstanceId *string `json:"DrdsInstanceId,omitempty" xml:"DrdsInstanceId,omitempty"`
	// The resource for which you want to configure parameters. Valid values:
	//
	// 	- **INSTANCE**: Configure parameters for the instance.
	//
	// 	- **DB**: Configure parameters for the databases of the instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// DB
	ParamLevel *string `json:"ParamLevel,omitempty" xml:"ParamLevel,omitempty"`
	// The ID of the region in which the PolarDB-X 1.0 instance is located.
	//
	// example:
	//
	// cn-hangzhou
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
}

func (s SetupDrdsParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s SetupDrdsParamsRequest) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequest) GetData() []*SetupDrdsParamsRequestData {
	return s.Data
}

func (s *SetupDrdsParamsRequest) GetDrdsInstanceId() *string {
	return s.DrdsInstanceId
}

func (s *SetupDrdsParamsRequest) GetParamLevel() *string {
	return s.ParamLevel
}

func (s *SetupDrdsParamsRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *SetupDrdsParamsRequest) SetData(v []*SetupDrdsParamsRequestData) *SetupDrdsParamsRequest {
	s.Data = v
	return s
}

func (s *SetupDrdsParamsRequest) SetDrdsInstanceId(v string) *SetupDrdsParamsRequest {
	s.DrdsInstanceId = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetParamLevel(v string) *SetupDrdsParamsRequest {
	s.ParamLevel = &v
	return s
}

func (s *SetupDrdsParamsRequest) SetRegionId(v string) *SetupDrdsParamsRequest {
	s.RegionId = &v
	return s
}

func (s *SetupDrdsParamsRequest) Validate() error {
	return dara.Validate(s)
}

type SetupDrdsParamsRequestData struct {
	// The name of the parameter that you want to configure for a database.
	//
	// example:
	//
	// test_db
	DbName *string `json:"DbName,omitempty" xml:"DbName,omitempty"`
	// The valid values of the parameter.
	//
	// example:
	//
	// [true|false]
	ParamRanges *string `json:"ParamRanges,omitempty" xml:"ParamRanges,omitempty"`
	// The type of the parameter that you want to configure. Valid values:
	//
	// 	- **ATOM**: the configuration item in the layer-3 data source.
	//
	// 	- **CONFIG**: the configuration item in ConfigServer.
	//
	// 	- **DIAMOND**: the configuration item in Diamond.
	//
	// example:
	//
	// ATOM
	ParamType *string `json:"ParamType,omitempty" xml:"ParamType,omitempty"`
	// The value of parameter that you want to configure.
	//
	// example:
	//
	// true
	ParamValue *string `json:"ParamValue,omitempty" xml:"ParamValue,omitempty"`
	// The name of the parameter that you want to configure.
	//
	// example:
	//
	// FORBID_EXECUTE_DML_ALL
	ParamVariableName *string `json:"ParamVariableName,omitempty" xml:"ParamVariableName,omitempty"`
}

func (s SetupDrdsParamsRequestData) String() string {
	return dara.Prettify(s)
}

func (s SetupDrdsParamsRequestData) GoString() string {
	return s.String()
}

func (s *SetupDrdsParamsRequestData) GetDbName() *string {
	return s.DbName
}

func (s *SetupDrdsParamsRequestData) GetParamRanges() *string {
	return s.ParamRanges
}

func (s *SetupDrdsParamsRequestData) GetParamType() *string {
	return s.ParamType
}

func (s *SetupDrdsParamsRequestData) GetParamValue() *string {
	return s.ParamValue
}

func (s *SetupDrdsParamsRequestData) GetParamVariableName() *string {
	return s.ParamVariableName
}

func (s *SetupDrdsParamsRequestData) SetDbName(v string) *SetupDrdsParamsRequestData {
	s.DbName = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamRanges(v string) *SetupDrdsParamsRequestData {
	s.ParamRanges = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamType(v string) *SetupDrdsParamsRequestData {
	s.ParamType = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamValue(v string) *SetupDrdsParamsRequestData {
	s.ParamValue = &v
	return s
}

func (s *SetupDrdsParamsRequestData) SetParamVariableName(v string) *SetupDrdsParamsRequestData {
	s.ParamVariableName = &v
	return s
}

func (s *SetupDrdsParamsRequestData) Validate() error {
	return dara.Validate(s)
}
