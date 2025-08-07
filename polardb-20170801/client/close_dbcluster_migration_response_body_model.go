// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseDBClusterMigrationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CloseDBClusterMigrationResponseBody
	GetRequestId() *string
}

type CloseDBClusterMigrationResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3AA69096-757C-4647-B36C-29EBC2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CloseDBClusterMigrationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseDBClusterMigrationResponseBody) GoString() string {
	return s.String()
}

func (s *CloseDBClusterMigrationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseDBClusterMigrationResponseBody) SetRequestId(v string) *CloseDBClusterMigrationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseDBClusterMigrationResponseBody) Validate() error {
	return dara.Validate(s)
}
