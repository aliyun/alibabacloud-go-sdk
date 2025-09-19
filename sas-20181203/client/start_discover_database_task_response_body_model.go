// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDiscoverDatabaseTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMark(v string) *StartDiscoverDatabaseTaskResponseBody
	GetCreateMark() *string
	SetRequestId(v string) *StartDiscoverDatabaseTaskResponseBody
	GetRequestId() *string
}

type StartDiscoverDatabaseTaskResponseBody struct {
	// The ID of the scan task.
	//
	// example:
	//
	// 48bced6d-2aee-4fa2-9aba-b846b77b****
	CreateMark *string `json:"CreateMark,omitempty" xml:"CreateMark,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F46921AF-CC55-5971-92C9-7E09E160****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDiscoverDatabaseTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDiscoverDatabaseTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StartDiscoverDatabaseTaskResponseBody) GetCreateMark() *string {
	return s.CreateMark
}

func (s *StartDiscoverDatabaseTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDiscoverDatabaseTaskResponseBody) SetCreateMark(v string) *StartDiscoverDatabaseTaskResponseBody {
	s.CreateMark = &v
	return s
}

func (s *StartDiscoverDatabaseTaskResponseBody) SetRequestId(v string) *StartDiscoverDatabaseTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDiscoverDatabaseTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
