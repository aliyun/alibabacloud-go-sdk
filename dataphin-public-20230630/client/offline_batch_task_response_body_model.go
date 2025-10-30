// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBatchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *OfflineBatchTaskResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *OfflineBatchTaskResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *OfflineBatchTaskResponseBody
	GetMessage() *string
	SetOfflineResult(v *OfflineBatchTaskResponseBodyOfflineResult) *OfflineBatchTaskResponseBody
	GetOfflineResult() *OfflineBatchTaskResponseBodyOfflineResult
	SetRequestId(v string) *OfflineBatchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *OfflineBatchTaskResponseBody
	GetSuccess() *bool
}

type OfflineBatchTaskResponseBody struct {
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
	Message       *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	OfflineResult *OfflineBatchTaskResponseBodyOfflineResult `json:"OfflineResult,omitempty" xml:"OfflineResult,omitempty" type:"Struct"`
	// example:
	//
	// 75DD06F8-1661-5A6E-B0A6-7E23133BDC60
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s OfflineBatchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineBatchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineBatchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *OfflineBatchTaskResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *OfflineBatchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *OfflineBatchTaskResponseBody) GetOfflineResult() *OfflineBatchTaskResponseBodyOfflineResult {
	return s.OfflineResult
}

func (s *OfflineBatchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineBatchTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *OfflineBatchTaskResponseBody) SetCode(v string) *OfflineBatchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *OfflineBatchTaskResponseBody) SetHttpStatusCode(v int32) *OfflineBatchTaskResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *OfflineBatchTaskResponseBody) SetMessage(v string) *OfflineBatchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *OfflineBatchTaskResponseBody) SetOfflineResult(v *OfflineBatchTaskResponseBodyOfflineResult) *OfflineBatchTaskResponseBody {
	s.OfflineResult = v
	return s
}

func (s *OfflineBatchTaskResponseBody) SetRequestId(v string) *OfflineBatchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineBatchTaskResponseBody) SetSuccess(v bool) *OfflineBatchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *OfflineBatchTaskResponseBody) Validate() error {
	if s.OfflineResult != nil {
		if err := s.OfflineResult.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OfflineBatchTaskResponseBodyOfflineResult struct {
	// example:
	//
	// 10211123
	PublishObjectId *int64 `json:"PublishObjectId,omitempty" xml:"PublishObjectId,omitempty"`
}

func (s OfflineBatchTaskResponseBodyOfflineResult) String() string {
	return dara.Prettify(s)
}

func (s OfflineBatchTaskResponseBodyOfflineResult) GoString() string {
	return s.String()
}

func (s *OfflineBatchTaskResponseBodyOfflineResult) GetPublishObjectId() *int64 {
	return s.PublishObjectId
}

func (s *OfflineBatchTaskResponseBodyOfflineResult) SetPublishObjectId(v int64) *OfflineBatchTaskResponseBodyOfflineResult {
	s.PublishObjectId = &v
	return s
}

func (s *OfflineBatchTaskResponseBodyOfflineResult) Validate() error {
	return dara.Validate(s)
}
