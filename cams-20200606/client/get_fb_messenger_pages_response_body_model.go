// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetFbMessengerPagesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessDeniedDetail(v string) *GetFbMessengerPagesResponseBody
	GetAccessDeniedDetail() *string
	SetCode(v string) *GetFbMessengerPagesResponseBody
	GetCode() *string
	SetData(v []*GetFbMessengerPagesResponseBodyData) *GetFbMessengerPagesResponseBody
	GetData() []*GetFbMessengerPagesResponseBodyData
	SetMessage(v string) *GetFbMessengerPagesResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetFbMessengerPagesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetFbMessengerPagesResponseBody
	GetSuccess() *bool
}

type GetFbMessengerPagesResponseBody struct {
	// example:
	//
	// {}
	AccessDeniedDetail *string `json:"AccessDeniedDetail,omitempty" xml:"AccessDeniedDetail,omitempty"`
	// example:
	//
	// OK
	Code *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data []*GetFbMessengerPagesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// sd2dsd-33*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetFbMessengerPagesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetFbMessengerPagesResponseBody) GoString() string {
	return s.String()
}

func (s *GetFbMessengerPagesResponseBody) GetAccessDeniedDetail() *string {
	return s.AccessDeniedDetail
}

func (s *GetFbMessengerPagesResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetFbMessengerPagesResponseBody) GetData() []*GetFbMessengerPagesResponseBodyData {
	return s.Data
}

func (s *GetFbMessengerPagesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetFbMessengerPagesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetFbMessengerPagesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetFbMessengerPagesResponseBody) SetAccessDeniedDetail(v string) *GetFbMessengerPagesResponseBody {
	s.AccessDeniedDetail = &v
	return s
}

func (s *GetFbMessengerPagesResponseBody) SetCode(v string) *GetFbMessengerPagesResponseBody {
	s.Code = &v
	return s
}

func (s *GetFbMessengerPagesResponseBody) SetData(v []*GetFbMessengerPagesResponseBodyData) *GetFbMessengerPagesResponseBody {
	s.Data = v
	return s
}

func (s *GetFbMessengerPagesResponseBody) SetMessage(v string) *GetFbMessengerPagesResponseBody {
	s.Message = &v
	return s
}

func (s *GetFbMessengerPagesResponseBody) SetRequestId(v string) *GetFbMessengerPagesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFbMessengerPagesResponseBody) SetSuccess(v bool) *GetFbMessengerPagesResponseBody {
	s.Success = &v
	return s
}

func (s *GetFbMessengerPagesResponseBody) Validate() error {
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

type GetFbMessengerPagesResponseBodyData struct {
	// example:
	//
	// 17433243434
	PageId *string `json:"PageId,omitempty" xml:"PageId,omitempty"`
	// example:
	//
	// iwhalecloud-2
	PageName *string `json:"PageName,omitempty" xml:"PageName,omitempty"`
}

func (s GetFbMessengerPagesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetFbMessengerPagesResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetFbMessengerPagesResponseBodyData) GetPageId() *string {
	return s.PageId
}

func (s *GetFbMessengerPagesResponseBodyData) GetPageName() *string {
	return s.PageName
}

func (s *GetFbMessengerPagesResponseBodyData) SetPageId(v string) *GetFbMessengerPagesResponseBodyData {
	s.PageId = &v
	return s
}

func (s *GetFbMessengerPagesResponseBodyData) SetPageName(v string) *GetFbMessengerPagesResponseBodyData {
	s.PageName = &v
	return s
}

func (s *GetFbMessengerPagesResponseBodyData) Validate() error {
	return dara.Validate(s)
}
