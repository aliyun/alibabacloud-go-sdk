// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNormalizationSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateNormalizationSchemaResponseBody
	GetRequestId() *string
}

type UpdateNormalizationSchemaResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNormalizationSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNormalizationSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNormalizationSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNormalizationSchemaResponseBody) SetRequestId(v string) *UpdateNormalizationSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNormalizationSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
