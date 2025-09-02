// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMetaCollectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateMetaCollectionResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateMetaCollectionResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *CreateMetaCollectionResponseBody
	GetHttpStatusCode() *string
	SetQualifiedName(v string) *CreateMetaCollectionResponseBody
	GetQualifiedName() *string
	SetRequestId(v string) *CreateMetaCollectionResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateMetaCollectionResponseBody
	GetSuccess() *string
}

type CreateMetaCollectionResponseBody struct {
	// The error code.
	//
	// example:
	//
	// NoPermission
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// The error message.
	//
	// example:
	//
	// The specified parameters are invalid.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The unique identifier of the collection.
	//
	// example:
	//
	// album.11111
	QualifiedName *string `json:"QualifiedName,omitempty" xml:"QualifiedName,omitempty"`
	// The request ID.
	//
	// example:
	//
	// E6F0DBDD-5AD****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateMetaCollectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMetaCollectionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMetaCollectionResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateMetaCollectionResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateMetaCollectionResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *CreateMetaCollectionResponseBody) GetQualifiedName() *string {
	return s.QualifiedName
}

func (s *CreateMetaCollectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMetaCollectionResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateMetaCollectionResponseBody) SetErrorCode(v string) *CreateMetaCollectionResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetErrorMessage(v string) *CreateMetaCollectionResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetHttpStatusCode(v string) *CreateMetaCollectionResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetQualifiedName(v string) *CreateMetaCollectionResponseBody {
	s.QualifiedName = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetRequestId(v string) *CreateMetaCollectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) SetSuccess(v string) *CreateMetaCollectionResponseBody {
	s.Success = &v
	return s
}

func (s *CreateMetaCollectionResponseBody) Validate() error {
	return dara.Validate(s)
}
