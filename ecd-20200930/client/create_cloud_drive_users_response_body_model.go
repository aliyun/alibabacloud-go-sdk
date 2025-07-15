// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudDriveUsersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCloudDriveUsersResponseBody
	GetRequestId() *string
}

type CreateCloudDriveUsersResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 05F0A7AE-17F1-53DF-BD99-ABF936FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCloudDriveUsersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudDriveUsersResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudDriveUsersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudDriveUsersResponseBody) SetRequestId(v string) *CreateCloudDriveUsersResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudDriveUsersResponseBody) Validate() error {
	return dara.Validate(s)
}
