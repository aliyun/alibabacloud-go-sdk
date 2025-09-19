// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartPreCheckDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCreateMark(v string) *StartPreCheckDatabaseResponseBody
	GetCreateMark() *string
	SetRequestId(v string) *StartPreCheckDatabaseResponseBody
	GetRequestId() *string
}

type StartPreCheckDatabaseResponseBody struct {
	// The ID of the database precheck task.
	//
	// example:
	//
	// t-0006d4pydyir6l1k****
	CreateMark *string `json:"CreateMark,omitempty" xml:"CreateMark,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// F6DC2DFF-AB3A-563A-8FC2-3D0D991E****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartPreCheckDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartPreCheckDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *StartPreCheckDatabaseResponseBody) GetCreateMark() *string {
	return s.CreateMark
}

func (s *StartPreCheckDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartPreCheckDatabaseResponseBody) SetCreateMark(v string) *StartPreCheckDatabaseResponseBody {
	s.CreateMark = &v
	return s
}

func (s *StartPreCheckDatabaseResponseBody) SetRequestId(v string) *StartPreCheckDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartPreCheckDatabaseResponseBody) Validate() error {
	return dara.Validate(s)
}
