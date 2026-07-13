// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetInstanceOssMountRamAuthorizeUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetCode() *string
	SetData(v *GetInstanceOssMountRamAuthorizeUrlResponseBodyData) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetData() *GetInstanceOssMountRamAuthorizeUrlResponseBodyData
	SetHttpStatusCode(v int32) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetInstanceOssMountRamAuthorizeUrlResponseBody
	GetSuccess() *bool
}

type GetInstanceOssMountRamAuthorizeUrlResponseBody struct {
	// example:
	//
	// SUCCESS
	Code *string                                             `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetInstanceOssMountRamAuthorizeUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// request-1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetInstanceOssMountRamAuthorizeUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceOssMountRamAuthorizeUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetData() *GetInstanceOssMountRamAuthorizeUrlResponseBodyData {
	return s.Data
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetCode(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetData(v *GetInstanceOssMountRamAuthorizeUrlResponseBodyData) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetHttpStatusCode(v int32) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetMessage(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetRequestId(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) SetSuccess(v bool) *GetInstanceOssMountRamAuthorizeUrlResponseBody {
	s.Success = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetInstanceOssMountRamAuthorizeUrlResponseBodyData struct {
	AuthorizeUrl *string `json:"AuthorizeUrl,omitempty" xml:"AuthorizeUrl,omitempty"`
}

func (s GetInstanceOssMountRamAuthorizeUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetInstanceOssMountRamAuthorizeUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBodyData) GetAuthorizeUrl() *string {
	return s.AuthorizeUrl
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBodyData) SetAuthorizeUrl(v string) *GetInstanceOssMountRamAuthorizeUrlResponseBodyData {
	s.AuthorizeUrl = &v
	return s
}

func (s *GetInstanceOssMountRamAuthorizeUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
