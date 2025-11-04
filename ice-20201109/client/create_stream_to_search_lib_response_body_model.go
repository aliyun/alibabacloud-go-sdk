// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStreamToSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateStreamToSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *CreateStreamToSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *CreateStreamToSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateStreamToSearchLibResponseBody
	GetSuccess() *string
}

type CreateStreamToSearchLibResponseBody struct {
	// The status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****20b48fb04483915d4f2cd8ac****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******3B-0E1A-586A-AC29-742247******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true: succeeded.
	//
	// 	- false: failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStreamToSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStreamToSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStreamToSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateStreamToSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CreateStreamToSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStreamToSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateStreamToSearchLibResponseBody) SetCode(v string) *CreateStreamToSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *CreateStreamToSearchLibResponseBody) SetMediaId(v string) *CreateStreamToSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *CreateStreamToSearchLibResponseBody) SetRequestId(v string) *CreateStreamToSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStreamToSearchLibResponseBody) SetSuccess(v string) *CreateStreamToSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStreamToSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
