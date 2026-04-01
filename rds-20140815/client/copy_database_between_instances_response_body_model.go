// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCopyDatabaseBetweenInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CopyDatabaseBetweenInstancesResponseBody
	GetRequestId() *string
}

type CopyDatabaseBetweenInstancesResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 803D11AF-C370-465B-AB46-CB3A642DC303
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CopyDatabaseBetweenInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CopyDatabaseBetweenInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *CopyDatabaseBetweenInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CopyDatabaseBetweenInstancesResponseBody) SetRequestId(v string) *CopyDatabaseBetweenInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CopyDatabaseBetweenInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}
