// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBroadcastVideosByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListBroadcastVideosByIdResponseBody
	GetCode() *string
	SetData(v []*BroadcastVideo) *ListBroadcastVideosByIdResponseBody
	GetData() []*BroadcastVideo
	SetHttpStatusCode(v int32) *ListBroadcastVideosByIdResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ListBroadcastVideosByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListBroadcastVideosByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListBroadcastVideosByIdResponseBody
	GetSuccess() *bool
}

type ListBroadcastVideosByIdResponseBody struct {
	// example:
	//
	// 200
	Code *string           `json:"code,omitempty" xml:"code,omitempty"`
	Data []*BroadcastVideo `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// SUCCESS
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// example:
	//
	// 7239F9E5-A4DB-55BA-B701-0CE47DBDB0A8
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ListBroadcastVideosByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListBroadcastVideosByIdResponseBody) GoString() string {
	return s.String()
}

func (s *ListBroadcastVideosByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListBroadcastVideosByIdResponseBody) GetData() []*BroadcastVideo {
	return s.Data
}

func (s *ListBroadcastVideosByIdResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListBroadcastVideosByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListBroadcastVideosByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListBroadcastVideosByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListBroadcastVideosByIdResponseBody) SetCode(v string) *ListBroadcastVideosByIdResponseBody {
	s.Code = &v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) SetData(v []*BroadcastVideo) *ListBroadcastVideosByIdResponseBody {
	s.Data = v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) SetHttpStatusCode(v int32) *ListBroadcastVideosByIdResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) SetMessage(v string) *ListBroadcastVideosByIdResponseBody {
	s.Message = &v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) SetRequestId(v string) *ListBroadcastVideosByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) SetSuccess(v bool) *ListBroadcastVideosByIdResponseBody {
	s.Success = &v
	return s
}

func (s *ListBroadcastVideosByIdResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
