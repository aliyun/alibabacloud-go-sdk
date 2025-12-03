// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeployKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *CreateDeployKeyResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *CreateDeployKeyResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateDeployKeyResponseBody
	GetRequestId() *string
	SetResult(v *CreateDeployKeyResponseBodyResult) *CreateDeployKeyResponseBody
	GetResult() *CreateDeployKeyResponseBodyResult
	SetSuccess(v bool) *CreateDeployKeyResponseBody
	GetSuccess() *bool
}

type CreateDeployKeyResponseBody struct {
	// example:
	//
	// SYSTEM_UNKNOWN_ERROR
	ErrorCode *string `json:"errorCode,omitempty" xml:"errorCode,omitempty"`
	// example:
	//
	// ”“
	ErrorMessage *string `json:"errorMessage,omitempty" xml:"errorMessage,omitempty"`
	// example:
	//
	// ASSDS-ASSASX-XSAXSA-XSAXSAXS
	RequestId *string                            `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Result    *CreateDeployKeyResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s CreateDeployKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeployKeyResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *CreateDeployKeyResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateDeployKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeployKeyResponseBody) GetResult() *CreateDeployKeyResponseBodyResult {
	return s.Result
}

func (s *CreateDeployKeyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDeployKeyResponseBody) SetErrorCode(v string) *CreateDeployKeyResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *CreateDeployKeyResponseBody) SetErrorMessage(v string) *CreateDeployKeyResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateDeployKeyResponseBody) SetRequestId(v string) *CreateDeployKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeployKeyResponseBody) SetResult(v *CreateDeployKeyResponseBodyResult) *CreateDeployKeyResponseBody {
	s.Result = v
	return s
}

func (s *CreateDeployKeyResponseBody) SetSuccess(v bool) *CreateDeployKeyResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDeployKeyResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDeployKeyResponseBodyResult struct {
	// example:
	//
	// 2022-03-18 14:24:54
	CreatedAt *string `json:"createdAt,omitempty" xml:"createdAt,omitempty"`
	// example:
	//
	// xx:xx:xx:xx
	Fingerprint *string `json:"fingerprint,omitempty" xml:"fingerprint,omitempty"`
	// example:
	//
	// 502385
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// ""
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
	Title *string `json:"title,omitempty" xml:"title,omitempty"`
}

func (s CreateDeployKeyResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateDeployKeyResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateDeployKeyResponseBodyResult) GetCreatedAt() *string {
	return s.CreatedAt
}

func (s *CreateDeployKeyResponseBodyResult) GetFingerprint() *string {
	return s.Fingerprint
}

func (s *CreateDeployKeyResponseBodyResult) GetId() *int64 {
	return s.Id
}

func (s *CreateDeployKeyResponseBodyResult) GetKey() *string {
	return s.Key
}

func (s *CreateDeployKeyResponseBodyResult) GetTitle() *string {
	return s.Title
}

func (s *CreateDeployKeyResponseBodyResult) SetCreatedAt(v string) *CreateDeployKeyResponseBodyResult {
	s.CreatedAt = &v
	return s
}

func (s *CreateDeployKeyResponseBodyResult) SetFingerprint(v string) *CreateDeployKeyResponseBodyResult {
	s.Fingerprint = &v
	return s
}

func (s *CreateDeployKeyResponseBodyResult) SetId(v int64) *CreateDeployKeyResponseBodyResult {
	s.Id = &v
	return s
}

func (s *CreateDeployKeyResponseBodyResult) SetKey(v string) *CreateDeployKeyResponseBodyResult {
	s.Key = &v
	return s
}

func (s *CreateDeployKeyResponseBodyResult) SetTitle(v string) *CreateDeployKeyResponseBodyResult {
	s.Title = &v
	return s
}

func (s *CreateDeployKeyResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
