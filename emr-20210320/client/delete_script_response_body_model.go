// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteScriptResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteScriptResponseBody
	GetRequestId() *string
}

type DeleteScriptResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// DD6B1B2A-5837-5237-ABE4-FF0C8944****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteScriptResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteScriptResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteScriptResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteScriptResponseBody) SetRequestId(v string) *DeleteScriptResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteScriptResponseBody) Validate() error {
	return dara.Validate(s)
}
