// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBNodesParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBNodeIds(v []*DescribeDBNodesParametersResponseBodyDBNodeIds) *DescribeDBNodesParametersResponseBody
	GetDBNodeIds() []*DescribeDBNodesParametersResponseBodyDBNodeIds
	SetDBType(v string) *DescribeDBNodesParametersResponseBody
	GetDBType() *string
	SetDBVersion(v string) *DescribeDBNodesParametersResponseBody
	GetDBVersion() *string
	SetEngine(v string) *DescribeDBNodesParametersResponseBody
	GetEngine() *string
	SetRequestId(v string) *DescribeDBNodesParametersResponseBody
	GetRequestId() *string
}

type DescribeDBNodesParametersResponseBody struct {
	// The list of node IDs.
	DBNodeIds []*DescribeDBNodesParametersResponseBodyDBNodeIds `json:"DBNodeIds,omitempty" xml:"DBNodeIds,omitempty" type:"Repeated"`
	// The database type. The value is fixed as **MySQL**.
	//
	// example:
	//
	// MySQL
	DBType *string `json:"DBType,omitempty" xml:"DBType,omitempty"`
	// The MySQL version. Valid values:
	//
	// - **5.6**
	//
	// - **5.7**
	//
	// - **8.0**
	//
	// example:
	//
	// 5.6
	DBVersion *string `json:"DBVersion,omitempty" xml:"DBVersion,omitempty"`
	// The cluster engine.
	//
	// example:
	//
	// POLARDB
	Engine *string `json:"Engine,omitempty" xml:"Engine,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9B7BFB11-C077-4FE3-B051-F69CEB******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBNodesParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodesParametersResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBNodesParametersResponseBody) GetDBNodeIds() []*DescribeDBNodesParametersResponseBodyDBNodeIds {
	return s.DBNodeIds
}

func (s *DescribeDBNodesParametersResponseBody) GetDBType() *string {
	return s.DBType
}

func (s *DescribeDBNodesParametersResponseBody) GetDBVersion() *string {
	return s.DBVersion
}

func (s *DescribeDBNodesParametersResponseBody) GetEngine() *string {
	return s.Engine
}

func (s *DescribeDBNodesParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBNodesParametersResponseBody) SetDBNodeIds(v []*DescribeDBNodesParametersResponseBodyDBNodeIds) *DescribeDBNodesParametersResponseBody {
	s.DBNodeIds = v
	return s
}

func (s *DescribeDBNodesParametersResponseBody) SetDBType(v string) *DescribeDBNodesParametersResponseBody {
	s.DBType = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBody) SetDBVersion(v string) *DescribeDBNodesParametersResponseBody {
	s.DBVersion = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBody) SetEngine(v string) *DescribeDBNodesParametersResponseBody {
	s.Engine = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBody) SetRequestId(v string) *DescribeDBNodesParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBody) Validate() error {
	if s.DBNodeIds != nil {
		for _, item := range s.DBNodeIds {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBNodesParametersResponseBodyDBNodeIds struct {
	// The node ID.
	//
	// example:
	//
	// pi-bp1r4qe3s534*****
	DBNodeId *string `json:"DBNodeId,omitempty" xml:"DBNodeId,omitempty"`
	// The list of running parameters for the current node.
	RunningParameters []*DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters `json:"RunningParameters,omitempty" xml:"RunningParameters,omitempty" type:"Repeated"`
}

func (s DescribeDBNodesParametersResponseBodyDBNodeIds) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodesParametersResponseBodyDBNodeIds) GoString() string {
	return s.String()
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIds) GetDBNodeId() *string {
	return s.DBNodeId
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIds) GetRunningParameters() []*DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	return s.RunningParameters
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIds) SetDBNodeId(v string) *DescribeDBNodesParametersResponseBodyDBNodeIds {
	s.DBNodeId = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIds) SetRunningParameters(v []*DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) *DescribeDBNodesParametersResponseBodyDBNodeIds {
	s.RunningParameters = v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIds) Validate() error {
	if s.RunningParameters != nil {
		for _, item := range s.RunningParameters {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters struct {
	// The value range of the parameter.
	//
	// example:
	//
	// [utf8|latin1|gbk|utf8mb4]
	CheckingCode *string `json:"CheckingCode,omitempty" xml:"CheckingCode,omitempty"`
	// The data type of the parameter value. Valid values:
	//
	// - **INT**: integer
	//
	// - **STRING**: string
	//
	// - **B**: byte
	//
	// example:
	//
	// INT
	DataType *string `json:"DataType,omitempty" xml:"DataType,omitempty"`
	// The default value of the parameter.
	//
	// example:
	//
	// utf8
	DefaultParameterValue *string `json:"DefaultParameterValue,omitempty" xml:"DefaultParameterValue,omitempty"`
	// The divisibility factor. For parameters of the integer and byte types, the parameter value must not only fall within the valid range but also be a multiple of the factor (which cannot be 0).
	//
	// example:
	//
	// 20
	Factor *string `json:"Factor,omitempty" xml:"Factor,omitempty"`
	// Specifies whether a restart is required for the parameter to take effect. Valid values:
	//
	// - **false**: no
	//
	// - **true**: yes
	//
	// example:
	//
	// true
	ForceRestart *bool `json:"ForceRestart,omitempty" xml:"ForceRestart,omitempty"`
	// Specifies whether the parameter is modifiable. Valid values:
	//
	// - **false**: not modifiable
	//
	// - **true**: modifiable
	//
	// example:
	//
	// true
	IsModifiable *bool `json:"IsModifiable,omitempty" xml:"IsModifiable,omitempty"`
	// Specifies whether the parameter is a global parameter. Valid values:
	//
	// - **0**: The parameter is a global parameter. Parameter modifications have a default application to other nodes and cannot be canceled.
	//
	// - **1**: The parameter is not a global parameter. You can specify the nodes to which modifications to this parameter are applied.
	//
	// example:
	//
	// 1
	IsNodeAvailable *string `json:"IsNodeAvailable,omitempty" xml:"IsNodeAvailable,omitempty"`
	// The parameter dependency.
	//
	// example:
	//
	// utf8
	ParamRelyRule *string `json:"ParamRelyRule,omitempty" xml:"ParamRelyRule,omitempty"`
	// The description of the parameter.
	//
	// example:
	//
	// The server\\"s default character set.
	ParameterDescription *string `json:"ParameterDescription,omitempty" xml:"ParameterDescription,omitempty"`
	// The name of the parameter.
	//
	// example:
	//
	// character_set_server
	ParameterName *string `json:"ParameterName,omitempty" xml:"ParameterName,omitempty"`
	// The status of the parameter. Valid values:
	//
	// - **normal**: normal
	//
	//
	//
	// - **modifying**: modifying
	//
	// example:
	//
	// normal
	ParameterStatus *string `json:"ParameterStatus,omitempty" xml:"ParameterStatus,omitempty"`
	// The value of the parameter.
	//
	// example:
	//
	// utf8
	ParameterValue *string `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GoString() string {
	return s.String()
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetCheckingCode() *string {
	return s.CheckingCode
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetDataType() *string {
	return s.DataType
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetDefaultParameterValue() *string {
	return s.DefaultParameterValue
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetFactor() *string {
	return s.Factor
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetForceRestart() *bool {
	return s.ForceRestart
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetIsModifiable() *bool {
	return s.IsModifiable
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetIsNodeAvailable() *string {
	return s.IsNodeAvailable
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetParamRelyRule() *string {
	return s.ParamRelyRule
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetParameterDescription() *string {
	return s.ParameterDescription
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetParameterName() *string {
	return s.ParameterName
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetParameterStatus() *string {
	return s.ParameterStatus
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetCheckingCode(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.CheckingCode = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetDataType(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.DataType = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetDefaultParameterValue(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.DefaultParameterValue = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetFactor(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.Factor = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetForceRestart(v bool) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ForceRestart = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetIsModifiable(v bool) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.IsModifiable = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetIsNodeAvailable(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.IsNodeAvailable = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetParamRelyRule(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ParamRelyRule = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetParameterDescription(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ParameterDescription = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetParameterName(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ParameterName = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetParameterStatus(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ParameterStatus = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) SetParameterValue(v string) *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters {
	s.ParameterValue = &v
	return s
}

func (s *DescribeDBNodesParametersResponseBodyDBNodeIdsRunningParameters) Validate() error {
	return dara.Validate(s)
}
