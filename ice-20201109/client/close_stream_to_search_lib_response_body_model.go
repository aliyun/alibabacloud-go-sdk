// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCloseStreamToSearchLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CloseStreamToSearchLibResponseBody
	GetCode() *string
	SetMediaId(v string) *CloseStreamToSearchLibResponseBody
	GetMediaId() *string
	SetRequestId(v string) *CloseStreamToSearchLibResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CloseStreamToSearchLibResponseBody
	GetSuccess() *string
}

type CloseStreamToSearchLibResponseBody struct {
	// The return code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the media asset.
	//
	// example:
	//
	// ****c469e944b5a856828dc2****
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
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

func (s CloseStreamToSearchLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CloseStreamToSearchLibResponseBody) GoString() string {
	return s.String()
}

func (s *CloseStreamToSearchLibResponseBody) GetCode() *string {
	return s.Code
}

func (s *CloseStreamToSearchLibResponseBody) GetMediaId() *string {
	return s.MediaId
}

func (s *CloseStreamToSearchLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CloseStreamToSearchLibResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CloseStreamToSearchLibResponseBody) SetCode(v string) *CloseStreamToSearchLibResponseBody {
	s.Code = &v
	return s
}

func (s *CloseStreamToSearchLibResponseBody) SetMediaId(v string) *CloseStreamToSearchLibResponseBody {
	s.MediaId = &v
	return s
}

func (s *CloseStreamToSearchLibResponseBody) SetRequestId(v string) *CloseStreamToSearchLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *CloseStreamToSearchLibResponseBody) SetSuccess(v string) *CloseStreamToSearchLibResponseBody {
	s.Success = &v
	return s
}

func (s *CloseStreamToSearchLibResponseBody) Validate() error {
	return dara.Validate(s)
}
