// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckDBNameAvailableResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CheckDBNameAvailableResponseBody
	GetRequestId() *string
}

type CheckDBNameAvailableResponseBody struct {
	// The request ID.
	//
	// >If only the request ID is returned, the database name is available. Otherwise, an error message is returned, indicating that the database name is duplicate or does not comply with naming conventions.
	//
	// example:
	//
	// 6EF82B07-28D2-48D1-B5D6-7E78FED277C7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CheckDBNameAvailableResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckDBNameAvailableResponseBody) GoString() string {
	return s.String()
}

func (s *CheckDBNameAvailableResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckDBNameAvailableResponseBody) SetRequestId(v string) *CheckDBNameAvailableResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckDBNameAvailableResponseBody) Validate() error {
	return dara.Validate(s)
}
