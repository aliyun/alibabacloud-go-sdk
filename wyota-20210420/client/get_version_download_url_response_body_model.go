// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVersionDownloadUrlResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetVersionDownloadUrlResponseBody
	GetCode() *string
	SetData(v *GetVersionDownloadUrlResponseBodyData) *GetVersionDownloadUrlResponseBody
	GetData() *GetVersionDownloadUrlResponseBodyData
	SetMessage(v string) *GetVersionDownloadUrlResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetVersionDownloadUrlResponseBody
	GetRequestId() *string
}

type GetVersionDownloadUrlResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetVersionDownloadUrlResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVersionDownloadUrlResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVersionDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetVersionDownloadUrlResponseBody) GetData() *GetVersionDownloadUrlResponseBodyData {
	return s.Data
}

func (s *GetVersionDownloadUrlResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetVersionDownloadUrlResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVersionDownloadUrlResponseBody) SetCode(v string) *GetVersionDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetData(v *GetVersionDownloadUrlResponseBodyData) *GetVersionDownloadUrlResponseBody {
	s.Data = v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetMessage(v string) *GetVersionDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) SetRequestId(v string) *GetVersionDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVersionDownloadUrlResponseBodyData struct {
	FullDownloadUrl *string `json:"FullDownloadUrl,omitempty" xml:"FullDownloadUrl,omitempty"`
}

func (s GetVersionDownloadUrlResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetVersionDownloadUrlResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetVersionDownloadUrlResponseBodyData) GetFullDownloadUrl() *string {
	return s.FullDownloadUrl
}

func (s *GetVersionDownloadUrlResponseBodyData) SetFullDownloadUrl(v string) *GetVersionDownloadUrlResponseBodyData {
	s.FullDownloadUrl = &v
	return s
}

func (s *GetVersionDownloadUrlResponseBodyData) Validate() error {
	return dara.Validate(s)
}
