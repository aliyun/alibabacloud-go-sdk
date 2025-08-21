// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpgradeEngineVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpgradeEngineVersionResponseBody
	GetRequestId() *string
	SetResult(v []*UpgradeEngineVersionResponseBodyResult) *UpgradeEngineVersionResponseBody
	GetResult() []*UpgradeEngineVersionResponseBodyResult
}

type UpgradeEngineVersionResponseBody struct {
	// The verification information.
	//
	// example:
	//
	// F99407AB-2FA9-489E-A259-40CF6DC*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The type of the error. Valid values:
	//
	// 	- clusterStatus: the health status of the cluster.
	//
	// 	- clusterConfigYml: Cluster YML File
	//
	// 	- clusterConfigPlugins: Cluster Configuration File
	//
	// 	- clusterResource: cluster resources
	//
	// 	- clusterSnapshot: cluster snapshot
	Result []*UpgradeEngineVersionResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s UpgradeEngineVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionResponseBody) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpgradeEngineVersionResponseBody) GetResult() []*UpgradeEngineVersionResponseBodyResult {
	return s.Result
}

func (s *UpgradeEngineVersionResponseBody) SetRequestId(v string) *UpgradeEngineVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpgradeEngineVersionResponseBody) SetResult(v []*UpgradeEngineVersionResponseBodyResult) *UpgradeEngineVersionResponseBody {
	s.Result = v
	return s
}

func (s *UpgradeEngineVersionResponseBody) Validate() error {
	return dara.Validate(s)
}

type UpgradeEngineVersionResponseBodyResult struct {
	// example:
	//
	// success
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The error message returned.
	ValidateResult []*UpgradeEngineVersionResponseBodyResultValidateResult `json:"validateResult,omitempty" xml:"validateResult,omitempty" type:"Repeated"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// checkClusterHealth
	ValidateType *string `json:"validateType,omitempty" xml:"validateType,omitempty"`
}

func (s UpgradeEngineVersionResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionResponseBodyResult) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBodyResult) GetStatus() *string {
	return s.Status
}

func (s *UpgradeEngineVersionResponseBodyResult) GetValidateResult() []*UpgradeEngineVersionResponseBodyResultValidateResult {
	return s.ValidateResult
}

func (s *UpgradeEngineVersionResponseBodyResult) GetValidateType() *string {
	return s.ValidateType
}

func (s *UpgradeEngineVersionResponseBodyResult) SetStatus(v string) *UpgradeEngineVersionResponseBodyResult {
	s.Status = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResult) SetValidateResult(v []*UpgradeEngineVersionResponseBodyResultValidateResult) *UpgradeEngineVersionResponseBodyResult {
	s.ValidateResult = v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResult) SetValidateType(v string) *UpgradeEngineVersionResponseBodyResult {
	s.ValidateType = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResult) Validate() error {
	return dara.Validate(s)
}

type UpgradeEngineVersionResponseBodyResultValidateResult struct {
	// example:
	//
	// ClusterStatusNotHealth
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// The cluster status is not health
	ErrorMsg *string `json:"errorMsg,omitempty" xml:"errorMsg,omitempty"`
	// The verification is passed. Valid values:
	//
	// 	- success: through
	//
	// 	- failed: failed
	//
	// example:
	//
	// clusterStatus
	ErrorType *string `json:"errorType,omitempty" xml:"errorType,omitempty"`
}

func (s UpgradeEngineVersionResponseBodyResultValidateResult) String() string {
	return dara.Prettify(s)
}

func (s UpgradeEngineVersionResponseBodyResultValidateResult) GoString() string {
	return s.String()
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) GetErrorMsg() *string {
	return s.ErrorMsg
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) GetErrorType() *string {
	return s.ErrorType
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorCode(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorCode = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorMsg(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorMsg = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) SetErrorType(v string) *UpgradeEngineVersionResponseBodyResultValidateResult {
	s.ErrorType = &v
	return s
}

func (s *UpgradeEngineVersionResponseBodyResultValidateResult) Validate() error {
	return dara.Validate(s)
}
