// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNacosMcpServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNacosMcpServerResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNacosMcpServerResponseBody
	GetRequestId() *string
}

type DeleteNacosMcpServerResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// D0DB055C-51F2-5BB2-82A6-CD8A3C6EE6BA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNacosMcpServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNacosMcpServerResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNacosMcpServerResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNacosMcpServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNacosMcpServerResponseBody) SetData(v bool) *DeleteNacosMcpServerResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNacosMcpServerResponseBody) SetRequestId(v string) *DeleteNacosMcpServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNacosMcpServerResponseBody) Validate() error {
	return dara.Validate(s)
}
