// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateIdentifyCredentialResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *CreateIdentifyCredentialResponseBody
	GetData() *string
	SetRequestId(v string) *CreateIdentifyCredentialResponseBody
	GetRequestId() *string
}

type CreateIdentifyCredentialResponseBody struct {
	// success true or false
	//
	// example:
	//
	// {   "HttpStatusCode": 200,   "Success": true }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateIdentifyCredentialResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateIdentifyCredentialResponseBody) GoString() string {
	return s.String()
}

func (s *CreateIdentifyCredentialResponseBody) GetData() *string {
	return s.Data
}

func (s *CreateIdentifyCredentialResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateIdentifyCredentialResponseBody) SetData(v string) *CreateIdentifyCredentialResponseBody {
	s.Data = &v
	return s
}

func (s *CreateIdentifyCredentialResponseBody) SetRequestId(v string) *CreateIdentifyCredentialResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateIdentifyCredentialResponseBody) Validate() error {
	return dara.Validate(s)
}
