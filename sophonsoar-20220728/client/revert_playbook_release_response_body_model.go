// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRevertPlaybookReleaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RevertPlaybookReleaseResponseBody
	GetRequestId() *string
}

type RevertPlaybookReleaseResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// B3FED5B9-190A-5952-93A4-24FBF0F0C573
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RevertPlaybookReleaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RevertPlaybookReleaseResponseBody) GoString() string {
	return s.String()
}

func (s *RevertPlaybookReleaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RevertPlaybookReleaseResponseBody) SetRequestId(v string) *RevertPlaybookReleaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *RevertPlaybookReleaseResponseBody) Validate() error {
	return dara.Validate(s)
}
