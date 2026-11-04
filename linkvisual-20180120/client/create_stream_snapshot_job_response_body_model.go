// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamSnapshotJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStreamSnapshotJobResponseBody
	GetCode() *string
	SetErrorMessage(v string) *CreateStreamSnapshotJobResponseBody
	GetErrorMessage() *string
	SetRequestId(v string) *CreateStreamSnapshotJobResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStreamSnapshotJobResponseBody
	GetSuccess() *bool
}

type CreateStreamSnapshotJobResponseBody struct {
	// The response code. Valid values:
	//
	// - **200**: The call was successful.
	//
	// - Other values: An error occurred. For more information, see [Error codes](https://help.aliyun.com/document_detail/145071.html).
	//
	// example:
	//
	// 403
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message returned if the call fails.
	//
	// example:
	//
	// Request auth failed.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// E55***B7-4***-4***-8***-D3******F565
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful. Valid values:
	//
	// - **true**: The call was successful.
	//
	// - **false**: The call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStreamSnapshotJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamSnapshotJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamSnapshotJobResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStreamSnapshotJobResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *CreateStreamSnapshotJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamSnapshotJobResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStreamSnapshotJobResponseBody) SetCode(v string) *CreateStreamSnapshotJobResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetErrorMessage(v string) *CreateStreamSnapshotJobResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetRequestId(v string) *CreateStreamSnapshotJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) SetSuccess(v bool) *CreateStreamSnapshotJobResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStreamSnapshotJobResponseBody) Validate() error {
	return dara.Validate(s)
}
