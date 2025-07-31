// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteOnlineTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteOnlineTestResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteOnlineTestResponseBody
	GetRequestId() *string
}

type DeleteOnlineTestResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteOnlineTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteOnlineTestResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteOnlineTestResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteOnlineTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteOnlineTestResponseBody) SetData(v bool) *DeleteOnlineTestResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteOnlineTestResponseBody) SetRequestId(v string) *DeleteOnlineTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteOnlineTestResponseBody) Validate() error {
	return dara.Validate(s)
}
