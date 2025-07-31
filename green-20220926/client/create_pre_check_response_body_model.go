// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreatePreCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]*bool) *CreatePreCheckResponseBody
	GetData() map[string]*bool
	SetRequestId(v string) *CreatePreCheckResponseBody
	GetRequestId() *string
}

type CreatePreCheckResponseBody struct {
	Data map[string]*bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreatePreCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreatePreCheckResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePreCheckResponseBody) GetData() map[string]*bool {
	return s.Data
}

func (s *CreatePreCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreatePreCheckResponseBody) SetData(v map[string]*bool) *CreatePreCheckResponseBody {
	s.Data = v
	return s
}

func (s *CreatePreCheckResponseBody) SetRequestId(v string) *CreatePreCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePreCheckResponseBody) Validate() error {
	return dara.Validate(s)
}
