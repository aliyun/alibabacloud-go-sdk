// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinueDBClusterMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ContinueDBClusterMigrationResponseBody
	GetRequestId() *string
}

type ContinueDBClusterMigrationResponseBody struct {
	// example:
	//
	// D2056BBE-FF76-5825-AB63-5CB1ABB46218
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinueDBClusterMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinueDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *ContinueDBClusterMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinueDBClusterMigrationResponseBody) SetRequestId(v string) *ContinueDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinueDBClusterMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
