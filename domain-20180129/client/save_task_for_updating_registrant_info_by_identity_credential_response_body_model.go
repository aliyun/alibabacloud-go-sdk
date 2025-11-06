// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody
	GetTaskNo() *string
}

type SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody struct {
	// example:
	//
	// EDC28FEC-6BE0-4583-95BC-test
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 880f1579-be51-4dd3-a69d-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByIdentityCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
