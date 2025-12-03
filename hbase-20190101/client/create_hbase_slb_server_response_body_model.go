// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHBaseSlbServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateHBaseSlbServerResponseBody
	GetRequestId() *string
}

type CreateHBaseSlbServerResponseBody struct {
	// example:
	//
	// 61FC5B21-87B0-41BC-9686-9DA395EB40B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateHBaseSlbServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateHBaseSlbServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateHBaseSlbServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateHBaseSlbServerResponseBody) SetRequestId(v string) *CreateHBaseSlbServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateHBaseSlbServerResponseBody) Validate() error {
	return dara.Validate(s)
}
