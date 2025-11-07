// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCloudauthstSceneResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateCloudauthstSceneResponseBody
	GetRequestId() *string
	SetResultObject(v bool) *CreateCloudauthstSceneResponseBody
	GetResultObject() *bool
}

type CreateCloudauthstSceneResponseBody struct {
	// ID of this request
	//
	// example:
	//
	// 5A6229C0-E156-48E4-B6EC-0F528BDF60D2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The result, indicating whether the operation was successful.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
}

func (s CreateCloudauthstSceneResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCloudauthstSceneResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCloudauthstSceneResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCloudauthstSceneResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *CreateCloudauthstSceneResponseBody) SetRequestId(v string) *CreateCloudauthstSceneResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCloudauthstSceneResponseBody) SetResultObject(v bool) *CreateCloudauthstSceneResponseBody {
	s.ResultObject = &v
	return s
}

func (s *CreateCloudauthstSceneResponseBody) Validate() error {
	return dara.Validate(s)
}
