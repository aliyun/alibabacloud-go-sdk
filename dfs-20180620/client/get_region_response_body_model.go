// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRegionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAvailableZones(v []*GetRegionResponseBodyAvailableZones) *GetRegionResponseBody
	GetAvailableZones() []*GetRegionResponseBodyAvailableZones
	SetRequestId(v string) *GetRegionResponseBody
	GetRequestId() *string
}

type GetRegionResponseBody struct {
	AvailableZones []*GetRegionResponseBodyAvailableZones `json:"AvailableZones,omitempty" xml:"AvailableZones,omitempty" type:"Repeated"`
	// example:
	//
	// AEC07154-5A4C-4B34-BB74-0893C6E9****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetRegionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetRegionResponseBody) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBody) GetAvailableZones() []*GetRegionResponseBodyAvailableZones {
	return s.AvailableZones
}

func (s *GetRegionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetRegionResponseBody) SetAvailableZones(v []*GetRegionResponseBodyAvailableZones) *GetRegionResponseBody {
	s.AvailableZones = v
	return s
}

func (s *GetRegionResponseBody) SetRequestId(v string) *GetRegionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetRegionResponseBody) Validate() error {
	if s.AvailableZones != nil {
		for _, item := range s.AvailableZones {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRegionResponseBodyAvailableZones struct {
	Options []*GetRegionResponseBodyAvailableZonesOptions `json:"Options,omitempty" xml:"Options,omitempty" type:"Repeated"`
	// example:
	//
	// cn-hangzhou-b
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s GetRegionResponseBodyAvailableZones) String() string {
	return dara.Prettify(s)
}

func (s GetRegionResponseBodyAvailableZones) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBodyAvailableZones) GetOptions() []*GetRegionResponseBodyAvailableZonesOptions {
	return s.Options
}

func (s *GetRegionResponseBodyAvailableZones) GetZoneId() *string {
	return s.ZoneId
}

func (s *GetRegionResponseBodyAvailableZones) SetOptions(v []*GetRegionResponseBodyAvailableZonesOptions) *GetRegionResponseBodyAvailableZones {
	s.Options = v
	return s
}

func (s *GetRegionResponseBodyAvailableZones) SetZoneId(v string) *GetRegionResponseBodyAvailableZones {
	s.ZoneId = &v
	return s
}

func (s *GetRegionResponseBodyAvailableZones) Validate() error {
	if s.Options != nil {
		for _, item := range s.Options {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetRegionResponseBodyAvailableZonesOptions struct {
	// example:
	//
	// HDFS
	ProtocolType *string `json:"ProtocolType,omitempty" xml:"ProtocolType,omitempty"`
	// example:
	//
	// STANDARD
	StorageType *string `json:"StorageType,omitempty" xml:"StorageType,omitempty"`
}

func (s GetRegionResponseBodyAvailableZonesOptions) String() string {
	return dara.Prettify(s)
}

func (s GetRegionResponseBodyAvailableZonesOptions) GoString() string {
	return s.String()
}

func (s *GetRegionResponseBodyAvailableZonesOptions) GetProtocolType() *string {
	return s.ProtocolType
}

func (s *GetRegionResponseBodyAvailableZonesOptions) GetStorageType() *string {
	return s.StorageType
}

func (s *GetRegionResponseBodyAvailableZonesOptions) SetProtocolType(v string) *GetRegionResponseBodyAvailableZonesOptions {
	s.ProtocolType = &v
	return s
}

func (s *GetRegionResponseBodyAvailableZonesOptions) SetStorageType(v string) *GetRegionResponseBodyAvailableZonesOptions {
	s.StorageType = &v
	return s
}

func (s *GetRegionResponseBodyAvailableZonesOptions) Validate() error {
	return dara.Validate(s)
}
