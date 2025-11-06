// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteContactTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteContactTemplatesResponseBody
	GetRequestId() *string
}

type DeleteContactTemplatesResponseBody struct {
	// example:
	//
	// 4D73432C-7600-4779-ACBB-C3B5CA145D32
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteContactTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteContactTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteContactTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteContactTemplatesResponseBody) SetRequestId(v string) *DeleteContactTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteContactTemplatesResponseBody) Validate() error {
	return dara.Validate(s)
}
