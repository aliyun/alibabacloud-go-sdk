// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddMediaSequencesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddMediaSequencesResponseBody
	GetRequestId() *string
}

type AddMediaSequencesResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddMediaSequencesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddMediaSequencesResponseBody) GoString() string {
	return s.String()
}

func (s *AddMediaSequencesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddMediaSequencesResponseBody) SetRequestId(v string) *AddMediaSequencesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddMediaSequencesResponseBody) Validate() error {
	return dara.Validate(s)
}
