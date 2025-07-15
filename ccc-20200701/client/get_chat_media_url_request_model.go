// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChatMediaUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *GetChatMediaUrlRequest
	GetInstanceId() *string
	SetMediaId(v string) *GetChatMediaUrlRequest
	GetMediaId() *string
	SetRequestId(v string) *GetChatMediaUrlRequest
	GetRequestId() *string
}

type GetChatMediaUrlRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// media id
	//
	// This parameter is required.
	//
	// example:
	//
	// $iAHNCNQCo3dhdgMGBAAFAAbaACOEAaQhIEeoAqpjjBl42N6o_kg7A88AAAGRIRRuBgTOACrxHgcACM8AAAGRIYJLBQ
	MediaId *string `json:"MediaId,omitempty" xml:"MediaId,omitempty"`
	// example:
	//
	// 8707EB29-BAED-4302-B999-40BA61877437
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetChatMediaUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetChatMediaUrlRequest) GoString() string {
	return s.String()
}

func (s *GetChatMediaUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetChatMediaUrlRequest) GetMediaId() *string {
	return s.MediaId
}

func (s *GetChatMediaUrlRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetChatMediaUrlRequest) SetInstanceId(v string) *GetChatMediaUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetChatMediaUrlRequest) SetMediaId(v string) *GetChatMediaUrlRequest {
	s.MediaId = &v
	return s
}

func (s *GetChatMediaUrlRequest) SetRequestId(v string) *GetChatMediaUrlRequest {
	s.RequestId = &v
	return s
}

func (s *GetChatMediaUrlRequest) Validate() error {
	return dara.Validate(s)
}
