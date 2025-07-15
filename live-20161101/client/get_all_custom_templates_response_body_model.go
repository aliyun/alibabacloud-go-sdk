// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllCustomTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCustomTemplates(v string) *GetAllCustomTemplatesResponseBody
	GetCustomTemplates() *string
	SetRequestId(v string) *GetAllCustomTemplatesResponseBody
	GetRequestId() *string
}

type GetAllCustomTemplatesResponseBody struct {
	// The template names and template configurations returned.
	//
	// example:
	//
	// [{"templateConfig": "{\\"cdesc\\":\\"H264\\",\\"scale\\":\\"[3:4]\\",\\"gop\\":\\"1\\",\\"bframes\\":\\"1\\",\\"height\\":\\"1080\\"}","templateName": "custom1"},{"templateConfig": "{\\"ar\\":\\"44100\\",\\"cdesc\\":\\"H264\\",\\"scale\\":\\"[3:4]\\",\\"gop\\":\\"1\\",\\"bframes\\":\\"1\\",\\"height\\":\\"1080\\"}","templateName": "cus"}]
	CustomTemplates *string `json:"CustomTemplates,omitempty" xml:"CustomTemplates,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 4791648Q-813C-6254-865C-0ED913661230
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAllCustomTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllCustomTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllCustomTemplatesResponseBody) GetCustomTemplates() *string {
	return s.CustomTemplates
}

func (s *GetAllCustomTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllCustomTemplatesResponseBody) SetCustomTemplates(v string) *GetAllCustomTemplatesResponseBody {
	s.CustomTemplates = &v
	return s
}

func (s *GetAllCustomTemplatesResponseBody) SetRequestId(v string) *GetAllCustomTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllCustomTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
