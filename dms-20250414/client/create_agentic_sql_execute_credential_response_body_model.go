// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAgenticSqlExecuteCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateAgenticSqlExecuteCredentialResponseBodyData) *CreateAgenticSqlExecuteCredentialResponseBody
	GetData() *CreateAgenticSqlExecuteCredentialResponseBodyData
	SetErrorCode(v string) *CreateAgenticSqlExecuteCredentialResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateAgenticSqlExecuteCredentialResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateAgenticSqlExecuteCredentialResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateAgenticSqlExecuteCredentialResponseBody
	GetSuccess() *bool
}

type CreateAgenticSqlExecuteCredentialResponseBody struct {
	// The temporary credential information.
	Data *CreateAgenticSqlExecuteCredentialResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error code.
	//
	// example:
	//
	// INVALID_PARAMETER
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// UnknownError
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 34E01EDD-6A16-4CF0-9541-C644D1BE01AA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// - **true**: The request is successful.
	//
	// - **false**: The request fails.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateAgenticSqlExecuteCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSqlExecuteCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) GetData() *CreateAgenticSqlExecuteCredentialResponseBodyData {
	return s.Data
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) SetData(v *CreateAgenticSqlExecuteCredentialResponseBodyData) *CreateAgenticSqlExecuteCredentialResponseBody {
	s.Data = v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) SetErrorCode(v string) *CreateAgenticSqlExecuteCredentialResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) SetErrorMessage(v string) *CreateAgenticSqlExecuteCredentialResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) SetRequestId(v string) *CreateAgenticSqlExecuteCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) SetSuccess(v bool) *CreateAgenticSqlExecuteCredentialResponseBody {
	s.Success = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateAgenticSqlExecuteCredentialResponseBodyData struct {
	// The AccessKey of the credential.
	//
	// example:
	//
	// ak_xxxxxxxx
	AccessKey *string `json:"AccessKey,omitempty" xml:"AccessKey,omitempty"`
	// The AccessKey secret of the credential. This value is returned only once when the credential is created. Keep it secure.
	//
	// example:
	//
	// xxxxxxxx
	AccessSecret *string `json:"AccessSecret,omitempty" xml:"AccessSecret,omitempty"`
	// The name of the instance.
	//
	// example:
	//
	// rm-bp1xxxxxxxx
	CatalogName *string `json:"CatalogName,omitempty" xml:"CatalogName,omitempty"`
	// The UUID of the instance.
	//
	// example:
	//
	// 9f8e7d6c-xxxx-xxxx-xxxx-1a2b3c4d5e6f
	CatalogUuid *string `json:"CatalogUuid,omitempty" xml:"CatalogUuid,omitempty"`
	// The UUID of the database. If the credential is at the instance level, this value is empty.
	//
	// example:
	//
	// 6c4e8a2b-xxxx-xxxx-xxxx-0f1d2e3c4b5a
	DatabaseUuid *string `json:"DatabaseUuid,omitempty" xml:"DatabaseUuid,omitempty"`
	// The expiration time of the credential. Format: yyyy-MM-dd HH:mm:ss.
	//
	// example:
	//
	// 2026-10-13 12:00:00
	ExpireTime *string `json:"ExpireTime,omitempty" xml:"ExpireTime,omitempty"`
	// The total number of connections supported by the credential.
	//
	// example:
	//
	// 500
	MaxConnectTimes *int64 `json:"MaxConnectTimes,omitempty" xml:"MaxConnectTimes,omitempty"`
	// The fully qualified name of the database. If the credential is at the instance level, this value is empty.
	//
	// example:
	//
	// test_db
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
}

func (s CreateAgenticSqlExecuteCredentialResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateAgenticSqlExecuteCredentialResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetAccessKey() *string {
	return s.AccessKey
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetAccessSecret() *string {
	return s.AccessSecret
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetCatalogName() *string {
	return s.CatalogName
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetCatalogUuid() *string {
	return s.CatalogUuid
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetDatabaseUuid() *string {
	return s.DatabaseUuid
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetExpireTime() *string {
	return s.ExpireTime
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetMaxConnectTimes() *int64 {
	return s.MaxConnectTimes
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetAccessKey(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.AccessKey = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetAccessSecret(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.AccessSecret = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetCatalogName(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.CatalogName = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetCatalogUuid(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.CatalogUuid = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetDatabaseUuid(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.DatabaseUuid = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetExpireTime(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.ExpireTime = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetMaxConnectTimes(v int64) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.MaxConnectTimes = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) SetQualifiedName(v string) *CreateAgenticSqlExecuteCredentialResponseBodyData {
	s.QualifiedName = &v
	return s
}

func (s *CreateAgenticSqlExecuteCredentialResponseBodyData) Validate() error {
	return dara.Validate(s)
}
