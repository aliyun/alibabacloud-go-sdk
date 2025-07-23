// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPauseInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *PauseInstanceResponseBody
	GetRequestId() *string
}

type PauseInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 4C467B38-3910-447D-87BC-AC049366F216
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PauseInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PauseInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *PauseInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PauseInstanceResponseBody) SetRequestId(v string) *PauseInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *PauseInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
