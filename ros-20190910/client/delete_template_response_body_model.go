// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteTemplateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteTemplateResponseBody
	GetRequestId() *string
}

type DeleteTemplateResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8C5D90E1-66B6-496C-9371-3807F8DA80A8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteTemplateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteTemplateResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteTemplateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteTemplateResponseBody) SetRequestId(v string) *DeleteTemplateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteTemplateResponseBody) Validate() error {
	return dara.Validate(s)
}
