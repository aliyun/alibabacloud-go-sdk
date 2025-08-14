// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSmartJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSmartJobResponseBody
	GetRequestId() *string
}

type DeleteSmartJobResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteSmartJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSmartJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSmartJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSmartJobResponseBody) SetRequestId(v string) *DeleteSmartJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSmartJobResponseBody) Validate() error {
	return dara.Validate(s)
}
