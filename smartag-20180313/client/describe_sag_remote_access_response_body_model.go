// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagRemoteAccessResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSagRemoteAccessResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeSagRemoteAccessResponseBody
	GetMessage() *string
	SetRemoteAccesses(v *DescribeSagRemoteAccessResponseBodyRemoteAccesses) *DescribeSagRemoteAccessResponseBody
	GetRemoteAccesses() *DescribeSagRemoteAccessResponseBodyRemoteAccesses
	SetRequestId(v string) *DescribeSagRemoteAccessResponseBody
	GetRequestId() *string
	SetSmartAGId(v string) *DescribeSagRemoteAccessResponseBody
	GetSmartAGId() *string
	SetSuccess(v bool) *DescribeSagRemoteAccessResponseBody
	GetSuccess() *bool
}

type DescribeSagRemoteAccessResponseBody struct {
	// The error code. The 200 error code indicates that the query task is successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message. The Successful error message indicates that the query task is successful.
	//
	// example:
	//
	// successful
	Message        *string                                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RemoteAccesses *DescribeSagRemoteAccessResponseBodyRemoteAccesses `json:"RemoteAccesses,omitempty" xml:"RemoteAccesses,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// E38E950D-28A4-4C41-9428-A8908EC6AE5C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-1um5x5nwhilymw****
	SmartAGId *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
	// Indicates whether the query task is successful. Valid values:
	//
	// 	- **true**: The query task is successful.
	//
	// 	- **false**: The query task failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSagRemoteAccessResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRemoteAccessResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagRemoteAccessResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSagRemoteAccessResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSagRemoteAccessResponseBody) GetRemoteAccesses() *DescribeSagRemoteAccessResponseBodyRemoteAccesses {
	return s.RemoteAccesses
}

func (s *DescribeSagRemoteAccessResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagRemoteAccessResponseBody) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeSagRemoteAccessResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSagRemoteAccessResponseBody) SetCode(v string) *DescribeSagRemoteAccessResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) SetMessage(v string) *DescribeSagRemoteAccessResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) SetRemoteAccesses(v *DescribeSagRemoteAccessResponseBodyRemoteAccesses) *DescribeSagRemoteAccessResponseBody {
	s.RemoteAccesses = v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) SetRequestId(v string) *DescribeSagRemoteAccessResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) SetSmartAGId(v string) *DescribeSagRemoteAccessResponseBody {
	s.SmartAGId = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) SetSuccess(v bool) *DescribeSagRemoteAccessResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBody) Validate() error {
	if s.RemoteAccesses != nil {
		if err := s.RemoteAccesses.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSagRemoteAccessResponseBodyRemoteAccesses struct {
	RemoteAccess []*DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess `json:"RemoteAccess,omitempty" xml:"RemoteAccess,omitempty" type:"Repeated"`
}

func (s DescribeSagRemoteAccessResponseBodyRemoteAccesses) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRemoteAccessResponseBodyRemoteAccesses) GoString() string {
	return s.String()
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccesses) GetRemoteAccess() []*DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess {
	return s.RemoteAccess
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccesses) SetRemoteAccess(v []*DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) *DescribeSagRemoteAccessResponseBodyRemoteAccesses {
	s.RemoteAccess = v
	return s
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccesses) Validate() error {
	if s.RemoteAccess != nil {
		for _, item := range s.RemoteAccess {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess struct {
	RemoteAccessIp *string `json:"RemoteAccessIp,omitempty" xml:"RemoteAccessIp,omitempty"`
	SerialNumber   *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
}

func (s DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) GoString() string {
	return s.String()
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) GetRemoteAccessIp() *string {
	return s.RemoteAccessIp
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) SetRemoteAccessIp(v string) *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess {
	s.RemoteAccessIp = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) SetSerialNumber(v string) *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess {
	s.SerialNumber = &v
	return s
}

func (s *DescribeSagRemoteAccessResponseBodyRemoteAccessesRemoteAccess) Validate() error {
	return dara.Validate(s)
}
