// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitBatchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *SubmitBatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *SubmitBatchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitBatchTaskResponseBody
	GetRequestId() *string
	SetSubmitResult(v *SubmitBatchTaskResponseBodySubmitResult) *SubmitBatchTaskResponseBody
	GetSubmitResult() *SubmitBatchTaskResponseBodySubmitResult
	SetSuccess(v bool) *SubmitBatchTaskResponseBody
	GetSuccess() *bool
}

type SubmitBatchTaskResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId    *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	SubmitResult *SubmitBatchTaskResponseBodySubmitResult `json:"SubmitResult,omitempty" xml:"SubmitResult,omitempty" type:"Struct"`
	Success      *bool                                    `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitBatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *SubmitBatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitBatchTaskResponseBody) GetSubmitResult() *SubmitBatchTaskResponseBodySubmitResult {
	return s.SubmitResult
}

func (s *SubmitBatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitBatchTaskResponseBody) SetCode(v string) *SubmitBatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitBatchTaskResponseBody) SetHttpStatusCode(v int32) *SubmitBatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *SubmitBatchTaskResponseBody) SetMessage(v string) *SubmitBatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitBatchTaskResponseBody) SetRequestId(v string) *SubmitBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitBatchTaskResponseBody) SetSubmitResult(v *SubmitBatchTaskResponseBodySubmitResult) *SubmitBatchTaskResponseBody {
	s.SubmitResult = v
	return s
}

func (s *SubmitBatchTaskResponseBody) SetSuccess(v bool) *SubmitBatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitBatchTaskResponseBody) Validate() error {
	if s.SubmitResult != nil {
		if err := s.SubmitResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitBatchTaskResponseBodySubmitResult struct {
	// example:
	//
	// n_12113111
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// example:
	//
	// 1010312911
	SubmitId *int64 `json:"SubmitId,omitempty" xml:"SubmitId,omitempty"`
}

func (s SubmitBatchTaskResponseBodySubmitResult) String() string {
	return dara.Prettify(s)
}

func (s SubmitBatchTaskResponseBodySubmitResult) GoString() string {
	return s.String()
}

func (s *SubmitBatchTaskResponseBodySubmitResult) GetNodeId() *string {
	return s.NodeId
}

func (s *SubmitBatchTaskResponseBodySubmitResult) GetSubmitId() *int64 {
	return s.SubmitId
}

func (s *SubmitBatchTaskResponseBodySubmitResult) SetNodeId(v string) *SubmitBatchTaskResponseBodySubmitResult {
	s.NodeId = &v
	return s
}

func (s *SubmitBatchTaskResponseBodySubmitResult) SetSubmitId(v int64) *SubmitBatchTaskResponseBodySubmitResult {
	s.SubmitId = &v
	return s
}

func (s *SubmitBatchTaskResponseBodySubmitResult) Validate() error {
	return dara.Validate(s)
}
