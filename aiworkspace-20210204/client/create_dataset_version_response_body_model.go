// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetVersionResponseBody
	GetRequestId() *string
	SetVersionName(v string) *CreateDatasetVersionResponseBody
	GetVersionName() *string
}

type CreateDatasetVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 41A847C8-3D12-5F24-8CE9-7F9EB2DA9ECD
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The name of the dataset version.
	//
	// example:
	//
	// v1
	VersionName *string `json:"VersionName,omitempty" xml:"VersionName,omitempty"`
}

func (s CreateDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetVersionResponseBody) GetVersionName() *string {
	return s.VersionName
}

func (s *CreateDatasetVersionResponseBody) SetRequestId(v string) *CreateDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) SetVersionName(v string) *CreateDatasetVersionResponseBody {
	s.VersionName = &v
	return s
}

func (s *CreateDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
