// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDedicatedClusterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDedicatedClusterId(v string) *ModifyDedicatedClusterResponseBody
	GetDedicatedClusterId() *string
	SetErrCode(v string) *ModifyDedicatedClusterResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDedicatedClusterResponseBody
	GetErrMessage() *string
	SetHttpStatusCode(v string) *ModifyDedicatedClusterResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *ModifyDedicatedClusterResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyDedicatedClusterResponseBody
	GetSuccess() *string
}

type ModifyDedicatedClusterResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// dtscluster_h3fl1cs217sx952
	DedicatedClusterId *string `json:"DedicatedClusterId,omitempty" xml:"DedicatedClusterId,omitempty"`
	// The error code returned if the request failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the request failed.
	//
	// example:
	//
	// The Value of Input Parameter %s is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 621BB4F8-3016-4FAA-8D5A-5D3163CC****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDedicatedClusterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDedicatedClusterResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDedicatedClusterResponseBody) GetDedicatedClusterId() *string {
	return s.DedicatedClusterId
}

func (s *ModifyDedicatedClusterResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDedicatedClusterResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDedicatedClusterResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *ModifyDedicatedClusterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDedicatedClusterResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyDedicatedClusterResponseBody) SetDedicatedClusterId(v string) *ModifyDedicatedClusterResponseBody {
	s.DedicatedClusterId = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) SetErrCode(v string) *ModifyDedicatedClusterResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) SetErrMessage(v string) *ModifyDedicatedClusterResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) SetHttpStatusCode(v string) *ModifyDedicatedClusterResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) SetRequestId(v string) *ModifyDedicatedClusterResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) SetSuccess(v string) *ModifyDedicatedClusterResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDedicatedClusterResponseBody) Validate() error {
	return dara.Validate(s)
}
