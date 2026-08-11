// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaLogResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateMetaLogResponseBody
	GetRequestId() *string
	SetResult(v bool) *UpdateMetaLogResponseBody
	GetResult() *bool
}

type UpdateMetaLogResponseBody struct {
	// The ID assigned by the backend to uniquely identify the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result.
	//
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s UpdateMetaLogResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaLogResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMetaLogResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMetaLogResponseBody) GetResult() *bool {
	return s.Result
}

func (s *UpdateMetaLogResponseBody) SetRequestId(v string) *UpdateMetaLogResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMetaLogResponseBody) SetResult(v bool) *UpdateMetaLogResponseBody {
	s.Result = &v
	return s
}

func (s *UpdateMetaLogResponseBody) Validate() error {
	return dara.Validate(s)
}
