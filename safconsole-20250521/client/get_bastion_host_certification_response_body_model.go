// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBastionHostCertificationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int64) *GetBastionHostCertificationResponseBody
	GetCode() *int64
	SetRequestId(v string) *GetBastionHostCertificationResponseBody
	GetRequestId() *string
	SetResultObject(v string) *GetBastionHostCertificationResponseBody
	GetResultObject() *string
	SetSuccess(v bool) *GetBastionHostCertificationResponseBody
	GetSuccess() *bool
}

type GetBastionHostCertificationResponseBody struct {
	// Status code. A return value of 200 indicates success.
	//
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Return result.
	//
	// example:
	//
	// xxx
	ResultObject *string `json:"ResultObject,omitempty" xml:"ResultObject,omitempty"`
	// Whether the call was successful.
	//
	// - **true**: Call succeeded.
	//
	// - **false**: Call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetBastionHostCertificationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetBastionHostCertificationResponseBody) GoString() string {
	return s.String()
}

func (s *GetBastionHostCertificationResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *GetBastionHostCertificationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetBastionHostCertificationResponseBody) GetResultObject() *string {
	return s.ResultObject
}

func (s *GetBastionHostCertificationResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetBastionHostCertificationResponseBody) SetCode(v int64) *GetBastionHostCertificationResponseBody {
	s.Code = &v
	return s
}

func (s *GetBastionHostCertificationResponseBody) SetRequestId(v string) *GetBastionHostCertificationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetBastionHostCertificationResponseBody) SetResultObject(v string) *GetBastionHostCertificationResponseBody {
	s.ResultObject = &v
	return s
}

func (s *GetBastionHostCertificationResponseBody) SetSuccess(v bool) *GetBastionHostCertificationResponseBody {
	s.Success = &v
	return s
}

func (s *GetBastionHostCertificationResponseBody) Validate() error {
	return dara.Validate(s)
}
