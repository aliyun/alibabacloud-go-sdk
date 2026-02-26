// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBClusterDynamoDBResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDBClusterDynamoDBResponseBody
	GetRequestId() *string
}

type DisableDBClusterDynamoDBResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBClusterDynamoDBResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBClusterDynamoDBResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBClusterDynamoDBResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBClusterDynamoDBResponseBody) SetRequestId(v string) *DisableDBClusterDynamoDBResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBClusterDynamoDBResponseBody) Validate() error {
	return dara.Validate(s)
}
