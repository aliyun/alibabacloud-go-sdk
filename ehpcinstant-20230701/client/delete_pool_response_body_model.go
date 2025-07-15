// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePoolResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeletePoolResponseBody
	GetRequestId() *string
}

type DeletePoolResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 896D338C-E4F4-41EC-A154-D605E5DE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeletePoolResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePoolResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePoolResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePoolResponseBody) SetRequestId(v string) *DeletePoolResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePoolResponseBody) Validate() error {
	return dara.Validate(s)
}
