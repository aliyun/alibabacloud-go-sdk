// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineLaboratoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OnlineLaboratoryResponseBody
	GetRequestId() *string
}

type OnlineLaboratoryResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8C27790E-CCA5-56BB-BA17-646295DEC0A2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OnlineLaboratoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OnlineLaboratoryResponseBody) GoString() string {
	return s.String()
}

func (s *OnlineLaboratoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OnlineLaboratoryResponseBody) SetRequestId(v string) *OnlineLaboratoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *OnlineLaboratoryResponseBody) Validate() error {
	return dara.Validate(s)
}
