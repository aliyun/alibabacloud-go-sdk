// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNormalizationSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateNormalizationSchemaResponseBody
	GetRequestId() *string
}

type CreateNormalizationSchemaResponseBody struct {
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNormalizationSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNormalizationSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNormalizationSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNormalizationSchemaResponseBody) SetRequestId(v string) *CreateNormalizationSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNormalizationSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
