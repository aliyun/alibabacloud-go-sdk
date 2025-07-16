// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLabelTableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLabelTableResponseBody
	GetRequestId() *string
}

type DeleteLabelTableResponseBody struct {
	// example:
	//
	// FFD39C0F-DD8D-51B2-864E-2842206DB0E8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLabelTableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLabelTableResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLabelTableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLabelTableResponseBody) SetRequestId(v string) *DeleteLabelTableResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLabelTableResponseBody) Validate() error {
	return dara.Validate(s)
}
