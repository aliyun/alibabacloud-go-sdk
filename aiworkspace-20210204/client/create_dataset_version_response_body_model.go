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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The dataset version name.
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
