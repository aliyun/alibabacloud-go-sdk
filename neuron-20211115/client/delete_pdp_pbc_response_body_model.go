// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePdpPbcResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeletePdpPbcResponseBody
	GetData() *bool
	SetRequestId(v string) *DeletePdpPbcResponseBody
	GetRequestId() *string
}

type DeletePdpPbcResponseBody struct {
	Data      *bool   `json:"data,omitempty" xml:"data,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s DeletePdpPbcResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeletePdpPbcResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePdpPbcResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeletePdpPbcResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeletePdpPbcResponseBody) SetData(v bool) *DeletePdpPbcResponseBody {
	s.Data = &v
	return s
}

func (s *DeletePdpPbcResponseBody) SetRequestId(v string) *DeletePdpPbcResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePdpPbcResponseBody) Validate() error {
	return dara.Validate(s)
}
