// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLogstashResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteLogstashResponseBody
	GetRequestId() *string
}

type DeleteLogstashResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 94B03BBA-A132-42C3-8367-0A0C1C45****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLogstashResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLogstashResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLogstashResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLogstashResponseBody) SetRequestId(v string) *DeleteLogstashResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLogstashResponseBody) Validate() error {
	return dara.Validate(s)
}
