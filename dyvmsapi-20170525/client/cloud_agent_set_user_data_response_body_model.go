// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentSetUserDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentSetUserDataResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentSetUserDataResponseBody
	GetCode() *string
	SetData(v *CloudAgentSetUserDataResponseBodyData) *CloudAgentSetUserDataResponseBody
	GetData() *CloudAgentSetUserDataResponseBodyData
	SetMessage(v string) *CloudAgentSetUserDataResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentSetUserDataResponseBody
	GetRequestId() *string
}

type CloudAgentSetUserDataResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentSetUserDataResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7BF47617-7851-48F7-A3A1-2021342A78E2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentSetUserDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentSetUserDataResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentSetUserDataResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentSetUserDataResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentSetUserDataResponseBody) GetData() *CloudAgentSetUserDataResponseBodyData {
	return s.Data
}

func (s *CloudAgentSetUserDataResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentSetUserDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentSetUserDataResponseBody) SetAccessDeniedDetail(v string) *CloudAgentSetUserDataResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentSetUserDataResponseBody) SetCode(v string) *CloudAgentSetUserDataResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentSetUserDataResponseBody) SetData(v *CloudAgentSetUserDataResponseBodyData) *CloudAgentSetUserDataResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentSetUserDataResponseBody) SetMessage(v string) *CloudAgentSetUserDataResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentSetUserDataResponseBody) SetRequestId(v string) *CloudAgentSetUserDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentSetUserDataResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentSetUserDataResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudAgentSetUserDataResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentSetUserDataResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentSetUserDataResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudAgentSetUserDataResponseBodyData) SetResult(v int64) *CloudAgentSetUserDataResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudAgentSetUserDataResponseBodyData) Validate() error {
	return dara.Validate(s)
}
