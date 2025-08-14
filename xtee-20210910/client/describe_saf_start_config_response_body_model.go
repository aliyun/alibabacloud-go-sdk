// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSafStartConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSafStartConfigResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeSafStartConfigResponseBodyResultObject) *DescribeSafStartConfigResponseBody
	GetResultObject() *DescribeSafStartConfigResponseBodyResultObject
}

type DescribeSafStartConfigResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returned object
	ResultObject *DescribeSafStartConfigResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
}

func (s DescribeSafStartConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSafStartConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSafStartConfigResponseBody) GetResultObject() *DescribeSafStartConfigResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeSafStartConfigResponseBody) SetRequestId(v string) *DescribeSafStartConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSafStartConfigResponseBody) SetResultObject(v *DescribeSafStartConfigResponseBodyResultObject) *DescribeSafStartConfigResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeSafStartConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSafStartConfigResponseBodyResultObject struct {
	// List of device types.
	DeviceTypes []*string `json:"deviceTypes,omitempty" xml:"deviceTypes,omitempty" type:"Repeated"`
	// Event codes.
	EventCodes []*string `json:"eventCodes,omitempty" xml:"eventCodes,omitempty" type:"Repeated"`
	// Configuration language details.
	Languages []*string `json:"languages,omitempty" xml:"languages,omitempty" type:"Repeated"`
	// Server region
	ServerRegions []*string `json:"serverRegions,omitempty" xml:"serverRegions,omitempty" type:"Repeated"`
}

func (s DescribeSafStartConfigResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeSafStartConfigResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeSafStartConfigResponseBodyResultObject) GetDeviceTypes() []*string {
	return s.DeviceTypes
}

func (s *DescribeSafStartConfigResponseBodyResultObject) GetEventCodes() []*string {
	return s.EventCodes
}

func (s *DescribeSafStartConfigResponseBodyResultObject) GetLanguages() []*string {
	return s.Languages
}

func (s *DescribeSafStartConfigResponseBodyResultObject) GetServerRegions() []*string {
	return s.ServerRegions
}

func (s *DescribeSafStartConfigResponseBodyResultObject) SetDeviceTypes(v []*string) *DescribeSafStartConfigResponseBodyResultObject {
	s.DeviceTypes = v
	return s
}

func (s *DescribeSafStartConfigResponseBodyResultObject) SetEventCodes(v []*string) *DescribeSafStartConfigResponseBodyResultObject {
	s.EventCodes = v
	return s
}

func (s *DescribeSafStartConfigResponseBodyResultObject) SetLanguages(v []*string) *DescribeSafStartConfigResponseBodyResultObject {
	s.Languages = v
	return s
}

func (s *DescribeSafStartConfigResponseBodyResultObject) SetServerRegions(v []*string) *DescribeSafStartConfigResponseBodyResultObject {
	s.ServerRegions = v
	return s
}

func (s *DescribeSafStartConfigResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
