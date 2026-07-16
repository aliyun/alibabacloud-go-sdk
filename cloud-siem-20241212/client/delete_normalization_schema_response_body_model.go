// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNormalizationSchemaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteNormalizationSchemaResponseBody
	GetRequestId() *string
}

type DeleteNormalizationSchemaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 6276D891-*****-55B2-87B9-74D413F7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNormalizationSchemaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNormalizationSchemaResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNormalizationSchemaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNormalizationSchemaResponseBody) SetRequestId(v string) *DeleteNormalizationSchemaResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNormalizationSchemaResponseBody) Validate() error {
	return dara.Validate(s)
}
