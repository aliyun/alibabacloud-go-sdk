// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadLinkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetUploadLinkResponseBody
	GetData() *string
	SetRequestId(v string) *GetUploadLinkResponseBody
	GetRequestId() *string
}

type GetUploadLinkResponseBody struct {
	// example:
	//
	// http://www.aliyuncs.com/test.mp3
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// E37E2B77-A664-5278-8219-23BBC0EB2A35
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUploadLinkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadLinkResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadLinkResponseBody) GetData() *string {
	return s.Data
}

func (s *GetUploadLinkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadLinkResponseBody) SetData(v string) *GetUploadLinkResponseBody {
	s.Data = &v
	return s
}

func (s *GetUploadLinkResponseBody) SetRequestId(v string) *GetUploadLinkResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadLinkResponseBody) Validate() error {
	return dara.Validate(s)
}
