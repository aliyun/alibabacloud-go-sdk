// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUploadContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *GetUploadContentResponseBody
	GetData() *string
	SetRequestId(v string) *GetUploadContentResponseBody
	GetRequestId() *string
}

type GetUploadContentResponseBody struct {
	// The response data.
	//
	// example:
	//
	// text
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetUploadContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUploadContentResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadContentResponseBody) GetData() *string {
	return s.Data
}

func (s *GetUploadContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUploadContentResponseBody) SetData(v string) *GetUploadContentResponseBody {
	s.Data = &v
	return s
}

func (s *GetUploadContentResponseBody) SetRequestId(v string) *GetUploadContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUploadContentResponseBody) Validate() error {
	return dara.Validate(s)
}
