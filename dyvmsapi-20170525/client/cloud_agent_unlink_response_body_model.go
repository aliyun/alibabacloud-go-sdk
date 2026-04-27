// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloudAgentUnlinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *CloudAgentUnlinkResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *CloudAgentUnlinkResponseBody
	GetCode() *string
	SetData(v *CloudAgentUnlinkResponseBodyData) *CloudAgentUnlinkResponseBody
	GetData() *CloudAgentUnlinkResponseBodyData
	SetMessage(v string) *CloudAgentUnlinkResponseBody
	GetMessage() *string
	SetRequestId(v string) *CloudAgentUnlinkResponseBody
	GetRequestId() *string
}

type CloudAgentUnlinkResponseBody struct {
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *CloudAgentUnlinkResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 6086693B-2250-17CE-A41F-06259AB6DB1B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloudAgentUnlinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentUnlinkResponseBody) GoString() string {
	return s.String()
}

func (s *CloudAgentUnlinkResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *CloudAgentUnlinkResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloudAgentUnlinkResponseBody) GetData() *CloudAgentUnlinkResponseBodyData {
	return s.Data
}

func (s *CloudAgentUnlinkResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CloudAgentUnlinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloudAgentUnlinkResponseBody) SetAccessDeniedDetail(v string) *CloudAgentUnlinkResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *CloudAgentUnlinkResponseBody) SetCode(v string) *CloudAgentUnlinkResponseBody {
	s.Code = &v
	return s
}

func (s *CloudAgentUnlinkResponseBody) SetData(v *CloudAgentUnlinkResponseBodyData) *CloudAgentUnlinkResponseBody {
	s.Data = v
	return s
}

func (s *CloudAgentUnlinkResponseBody) SetMessage(v string) *CloudAgentUnlinkResponseBody {
	s.Message = &v
	return s
}

func (s *CloudAgentUnlinkResponseBody) SetRequestId(v string) *CloudAgentUnlinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloudAgentUnlinkResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CloudAgentUnlinkResponseBodyData struct {
	// 结果
	//
	// example:
	//
	// 0
	Result *int64 `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s CloudAgentUnlinkResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CloudAgentUnlinkResponseBodyData) GoString() string {
	return s.String()
}

func (s *CloudAgentUnlinkResponseBodyData) GetResult() *int64 {
	return s.Result
}

func (s *CloudAgentUnlinkResponseBodyData) SetResult(v int64) *CloudAgentUnlinkResponseBodyData {
	s.Result = &v
	return s
}

func (s *CloudAgentUnlinkResponseBodyData) Validate() error {
	return dara.Validate(s)
}
