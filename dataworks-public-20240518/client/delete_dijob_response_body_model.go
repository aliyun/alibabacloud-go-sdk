// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDIJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDIJobResponseBody
	GetSuccess() *bool
}

type DeleteDIJobResponseBody struct {
	// The request ID. You can troubleshoot issues based on the ID.
	//
	// example:
	//
	// D33D4A51-5845-579A-B4BA-FAADD0F83D53
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// true
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDIJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDIJobResponseBody) SetRequestId(v string) *DeleteDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDIJobResponseBody) SetSuccess(v bool) *DeleteDIJobResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
