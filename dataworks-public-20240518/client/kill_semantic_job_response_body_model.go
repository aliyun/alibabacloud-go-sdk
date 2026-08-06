// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iKillSemanticJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *KillSemanticJobResponseBody
	GetData() *bool
	SetRequestId(v string) *KillSemanticJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *KillSemanticJobResponseBody
	GetSuccess() *bool
}

type KillSemanticJobResponseBody struct {
	// Indicates whether the stop request has been accepted by the executor. Even if true is returned, call GetSemanticJobDetail to query the final status.
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID. Used for locating logs and troubleshooting issues.
	//
	// example:
	//
	// 676271D6-53B4-57BE-89FA-72F7AE1418DF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s KillSemanticJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s KillSemanticJobResponseBody) GoString() string {
	return s.String()
}

func (s *KillSemanticJobResponseBody) GetData() *bool {
	return s.Data
}

func (s *KillSemanticJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *KillSemanticJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *KillSemanticJobResponseBody) SetData(v bool) *KillSemanticJobResponseBody {
	s.Data = &v
	return s
}

func (s *KillSemanticJobResponseBody) SetRequestId(v string) *KillSemanticJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *KillSemanticJobResponseBody) SetSuccess(v bool) *KillSemanticJobResponseBody {
	s.Success = &v
	return s
}

func (s *KillSemanticJobResponseBody) Validate() error {
	return dara.Validate(s)
}
