// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTotalStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNetworkOut(v int64) *GetTotalStatisResponseBody
	GetNetworkOut() *int64
	SetRequestId(v string) *GetTotalStatisResponseBody
	GetRequestId() *string
	SetStorageUtilization(v int64) *GetTotalStatisResponseBody
	GetStorageUtilization() *int64
	SetTranscodeDuration(v int64) *GetTotalStatisResponseBody
	GetTranscodeDuration() *int64
}

type GetTotalStatisResponseBody struct {
	NetworkOut         *int64  `json:"NetworkOut,omitempty" xml:"NetworkOut,omitempty"`
	RequestId          *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageUtilization *int64  `json:"StorageUtilization,omitempty" xml:"StorageUtilization,omitempty"`
	TranscodeDuration  *int64  `json:"TranscodeDuration,omitempty" xml:"TranscodeDuration,omitempty"`
}

func (s GetTotalStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTotalStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetTotalStatisResponseBody) GetNetworkOut() *int64 {
	return s.NetworkOut
}

func (s *GetTotalStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTotalStatisResponseBody) GetStorageUtilization() *int64 {
	return s.StorageUtilization
}

func (s *GetTotalStatisResponseBody) GetTranscodeDuration() *int64 {
	return s.TranscodeDuration
}

func (s *GetTotalStatisResponseBody) SetNetworkOut(v int64) *GetTotalStatisResponseBody {
	s.NetworkOut = &v
	return s
}

func (s *GetTotalStatisResponseBody) SetRequestId(v string) *GetTotalStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTotalStatisResponseBody) SetStorageUtilization(v int64) *GetTotalStatisResponseBody {
	s.StorageUtilization = &v
	return s
}

func (s *GetTotalStatisResponseBody) SetTranscodeDuration(v int64) *GetTotalStatisResponseBody {
	s.TranscodeDuration = &v
	return s
}

func (s *GetTotalStatisResponseBody) Validate() error {
	return dara.Validate(s)
}
