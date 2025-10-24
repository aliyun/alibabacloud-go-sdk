// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmarttagTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSmarttagTemplateResponseBody
	GetRequestId() *string
}

type DeleteSmarttagTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5F37036F-5267-43F1-AE47-10A18E840739
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmarttagTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmarttagTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmarttagTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmarttagTemplateResponseBody) SetRequestId(v string) *DeleteSmarttagTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmarttagTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
