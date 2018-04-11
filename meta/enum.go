// Trimmer SDK
//
// Copyright (c) 2017-2018 Alexander Eichhorn
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package meta

import (
	trimmer "trimmer.io/go-trimmer"
)

// Groups
const (
	NoMetadata     trimmer.MetaGroup = ""
	XmpMetadata    trimmer.MetaGroup = "xmp"
	ImageMetadata  trimmer.MetaGroup = "image"
	MusicMetadata  trimmer.MetaGroup = "music"
	MovieMetadata  trimmer.MetaGroup = "movie"
	SoundMetadata  trimmer.MetaGroup = "sound"
	CameraMetadata trimmer.MetaGroup = "camera"
	VfxMetadata    trimmer.MetaGroup = "vfx"
)

// Trimmer Metadata Paths
const (
	AssetCreateDate  trimmer.MetaPath = "trim:Asset/CreateDate"
	AssetIngestDate  trimmer.MetaPath = "trim:Asset/IngestDate"
	AssetType        trimmer.MetaPath = "trim:Asset/Type"
	AssetTitle       trimmer.MetaPath = "trim:Asset/Title"
	AssetKeywords    trimmer.MetaPath = "trim:Asset/Keywords"
	AssetDescription trimmer.MetaPath = "trim:Asset/Description"
	AssetNotes       trimmer.MetaPath = "trim:Asset/Notes"
	AssetCopyright   trimmer.MetaPath = "trim:Asset/Copyright"
	AssetLicense     trimmer.MetaPath = "trim:Asset/License"
	AssetUUID        trimmer.MetaPath = "trim:Asset/UUID"

	CameraId                 trimmer.MetaPath = "trim:Camera/Id"
	CameraIndex              trimmer.MetaPath = "trim:Camera/Index"
	CameraMake               trimmer.MetaPath = "trim:Camera/Make"
	CameraModel              trimmer.MetaPath = "trim:Camera/Model"
	CameraSoftware           trimmer.MetaPath = "trim:Camera/Software"
	CameraSerial             trimmer.MetaPath = "trim:Camera/Serial"
	CameraMakeDate           trimmer.MetaPath = "trim:Camera/MakeDate"
	CameraNotes              trimmer.MetaPath = "trim:Camera/Notes"
	CameraIdentifier         trimmer.MetaPath = "trim:Camera/Identifier"
	CameraLabel              trimmer.MetaPath = "trim:Camera/Label"
	CameraRecorderType       trimmer.MetaPath = "trim:Camera/RecorderType"
	CameraExposureIndex      trimmer.MetaPath = "trim:Camera/ExposureIndex"
	CameraExposureTime       trimmer.MetaPath = "trim:Camera/ExposureTime"
	CameraFrameRate          trimmer.MetaPath = "trim:Camera/FrameRate"
	CameraShutterSpeed       trimmer.MetaPath = "trim:Camera/ShutterSpeed"
	CameraShutterAngle       trimmer.MetaPath = "trim:Camera/ShutterAngle"
	CameraWhiteBalanceKelvin trimmer.MetaPath = "trim:Camera/WhiteBalanceKelvin"
	CameraWhiteBalanceTintCc trimmer.MetaPath = "trim:Camera/WhiteBalanceTintCc"
	CameraTemperature        trimmer.MetaPath = "trim:Camera/Temperature"
	CameraOrientation        trimmer.MetaPath = "trim:Camera/Orientation"
	CameraColorSpace         trimmer.MetaPath = "trim:Camera/ColorSpace"
	CameraLookName           trimmer.MetaPath = "trim:Camera/LookName"
	CameraLookID             trimmer.MetaPath = "trim:Camera/LookId"

	ContentPictureQuality  trimmer.MetaPath = "trim:Content/PictureQuality"
	ContentSoundQuality    trimmer.MetaPath = "trim:Content/SoundQuality"
	ContentSubtitleQuality trimmer.MetaPath = "trim:Content/SubtitleQuality"
	ContentLanguage        trimmer.MetaPath = "trim:Content/Language"
	ContentNotes           trimmer.MetaPath = "trim:Content/Notes"
	ContentKeywords        trimmer.MetaPath = "trim:Content/Keywords"
	ContentBrand           trimmer.MetaPath = "trim:Content/Brand"
	ContentStyle           trimmer.MetaPath = "trim:Content/Style"
	ContentProduct         trimmer.MetaPath = "trim:Content/Product"
	ContentEvent           trimmer.MetaPath = "trim:Content/Event"
	ContentProperty        trimmer.MetaPath = "trim:Content/Property"
	ContentPeople          trimmer.MetaPath = "trim:Content/People"
	ContentLocation        trimmer.MetaPath = "trim:Content/Location"

	CrewUnit             trimmer.MetaPath = "trim:Crew/Unit"
	CrewDirector         trimmer.MetaPath = "trim:Crew/Director"
	CrewCinematographer  trimmer.MetaPath = "trim:Crew/Cinematographer"
	CrewCameraOperator   trimmer.MetaPath = "trim:Crew/CameraOperator"
	CrewEditor           trimmer.MetaPath = "trim:Crew/Editor"
	CrewEngineer         trimmer.MetaPath = "trim:Crew/Engineer"
	CrewCreator          trimmer.MetaPath = "trim:Crew/Creator"
	CrewScreenwriter     trimmer.MetaPath = "trim:Crew/Screenwriter"
	CrewProducer         trimmer.MetaPath = "trim:Crew/Producer"
	CrewActor            trimmer.MetaPath = "trim:Crew/Actor"
	CrewArtDirector      trimmer.MetaPath = "trim:Crew/ArtDirector"
	CrewCreativeDirector trimmer.MetaPath = "trim:Crew/CreativeDirector"
	CrewCopywriter       trimmer.MetaPath = "trim:Crew/Copywriter"
	CrewColorist         trimmer.MetaPath = "trim:Crew/Colorist"
	CrewSoundDesigner    trimmer.MetaPath = "trim:Crew/SoundDesigner"
	CrewMusician         trimmer.MetaPath = "trim:Crew/Musician"
	CrewComposer         trimmer.MetaPath = "trim:Crew/Composer"
	CrewVfxArtist        trimmer.MetaPath = "trim:Crew/VfxArtist"
	CrewLegalAdvisor     trimmer.MetaPath = "trim:Crew/LegalAdvisor"

	DeviceId       trimmer.MetaPath = "trim:Device/Id"
	DeviceIndex    trimmer.MetaPath = "trim:Device/Index"
	DeviceMake     trimmer.MetaPath = "trim:Device/Make"
	DeviceModel    trimmer.MetaPath = "trim:Device/Model"
	DeviceSoftware trimmer.MetaPath = "trim:Device/Software"
	DeviceSerial   trimmer.MetaPath = "trim:Device/Serial"
	DeviceMakeDate trimmer.MetaPath = "trim:Device/MakeDate"
	DeviceNOtes    trimmer.MetaPath = "trim:Device/Notes"

	DistributionIdentifier          trimmer.MetaPath = "trim:Distribution/Identifier"
	DistributionWorkType            trimmer.MetaPath = "trim:Distribution/WorkType"
	DistributionTitle               trimmer.MetaPath = "trim:Distribution/Title"
	DistributionSeason              trimmer.MetaPath = "trim:Distribution/Season"
	DistributionEpisode             trimmer.MetaPath = "trim:Distribution/Episode"
	DistributionNetwork             trimmer.MetaPath = "trim:Distribution/Network"
	DistributionSynopsis            trimmer.MetaPath = "trim:Distribution/Synopsis"
	DistributionGenre               trimmer.MetaPath = "trim:Distribution/Genre"
	DistributionKeywords            trimmer.MetaPath = "trim:Distribution/Keywords"
	DistributionOriginLanguage      trimmer.MetaPath = "trim:Distribution/OriginLanguage"
	DistributionOriginCountry       trimmer.MetaPath = "trim:Distribution/OriginCountry"
	DistributionCompletionDate      trimmer.MetaPath = "trim:Distribution/CompletionDate"
	DistributionCopyrightDate       trimmer.MetaPath = "trim:Distribution/CopyrightDate"
	DistributionReleaseDate         trimmer.MetaPath = "trim:Distribution/ReleaseDate"
	DistributionHasCaptions         trimmer.MetaPath = "trim:Distribution/HasCaptions"
	DistributionHasSigning          trimmer.MetaPath = "trim:Distribution/HasSigning"
	DistributionHasProductPlacement trimmer.MetaPath = "trim:Distribution/HasProductPlacement"
	DistributionAudioLoudness       trimmer.MetaPath = "trim:Distribution/AudioLoudness"

	EncoderId       trimmer.MetaPath = "trim:Encoder/Id"
	EncoderIndex    trimmer.MetaPath = "trim:Encoder/Index"
	EncoderMake     trimmer.MetaPath = "trim:Encoder/Make"
	EncoderModel    trimmer.MetaPath = "trim:Encoder/Model"
	EncoderSoftware trimmer.MetaPath = "trim:Encoder/Software"
	EncoderSerial   trimmer.MetaPath = "trim:Encoder/Serial"
	EncoderMakeDate trimmer.MetaPath = "trim:Encoder/MakeDate"
	EncoderNotes    trimmer.MetaPath = "trim:Encoder/Notes"
	EncoderHost     trimmer.MetaPath = "trim:Encoder/Host"

	LensId            trimmer.MetaPath = "trim:Lens/Id"
	LensIndex         trimmer.MetaPath = "trim:Lens/Index"
	LensMake          trimmer.MetaPath = "trim:Lens/Make"
	LensModel         trimmer.MetaPath = "trim:Lens/Model"
	LensSoftware      trimmer.MetaPath = "trim:Lens/Software"
	LensSerial        trimmer.MetaPath = "trim:Lens/Serial"
	LensMakeDate      trimmer.MetaPath = "trim:Lens/MakeDate"
	LensNotes         trimmer.MetaPath = "trim:Lens/Notes"
	LensFocalLength   trimmer.MetaPath = "trim:Lens/FocalLength"
	LensSqueeze       trimmer.MetaPath = "trim:Lens/Squeeze"
	LensFilterDensity trimmer.MetaPath = "trim:Lens/FilterDensity"
	LensFilterType    trimmer.MetaPath = "trim:Lens/FilterType"

	LocationName            trimmer.MetaPath = "trim:Location/Name"
	LocationRoom            trimmer.MetaPath = "trim:Location/Room"
	LocationAddress         trimmer.MetaPath = "trim:Location/Address"
	LocationCity            trimmer.MetaPath = "trim:Location/City"
	LocationState           trimmer.MetaPath = "trim:Location/State"
	LocationRegion          trimmer.MetaPath = "trim:Location/Region"
	LocationPostCode        trimmer.MetaPath = "trim:Location/PostCode"
	LocationCountryCode     trimmer.MetaPath = "trim:Location/CountryCode"
	LocationCountryName     trimmer.MetaPath = "trim:Location/CountryName"
	LocationBody            trimmer.MetaPath = "trim:Location/Body"
	LocationGPSLatitude     trimmer.MetaPath = "trim:Location/GpsLatitude"
	LocationGPSLongitude    trimmer.MetaPath = "trim:Location/GpsLongitude"
	LocationGPSAltitude     trimmer.MetaPath = "trim:Location/GpsAltitude"
	LocationGPSSpeed        trimmer.MetaPath = "trim:Location/GpsSpeed"
	LocationGPSAccuracy     trimmer.MetaPath = "trim:Location/GpsAccuracy"
	LocationGPSImgDirection trimmer.MetaPath = "trim:Location/GpsImgDirection"
	LocationGPSDestBearing  trimmer.MetaPath = "trim:Location/GpsDestBearing"
	LocationGPSDestDistance trimmer.MetaPath = "trim:Location/GpsDestDistance"

	ProductionName      trimmer.MetaPath = "trim:Production/Name"
	ProductionCompany   trimmer.MetaPath = "trim:Production/Company"
	ProductionAgency    trimmer.MetaPath = "trim:Production/Agency"
	ProductionClient    trimmer.MetaPath = "trim:Production/Client"
	ProductionPostHouse trimmer.MetaPath = "trim:Production/PostHouse"
	ProductionVfxHouse  trimmer.MetaPath = "trim:Production/VfxHouse"

	ShotClipName     trimmer.MetaPath = "trim:Shot/ClipName"
	ShotSceneName    trimmer.MetaPath = "trim:Shot/Scene"
	ShotSetupName    trimmer.MetaPath = "trim:Shot/Setup"
	ShotReelName     trimmer.MetaPath = "trim:Shot/Reel"
	ShotTakeName     trimmer.MetaPath = "trim:Shot/Take"
	ShotAngle        trimmer.MetaPath = "trim:Shot/Angle"
	ShotDay          trimmer.MetaPath = "trim:Shot/Day"
	ShotSoundReel    trimmer.MetaPath = "trim:Shot/SoundReel"
	ShotSoundFile    trimmer.MetaPath = "trim:Shot/SoundFile"
	ShotIsCircled    trimmer.MetaPath = "trim:Shot/IsCircled"
	ShotIsIndoor     trimmer.MetaPath = "trim:Shot/IsIndoor"
	ShotIsOutdoor    trimmer.MetaPath = "trim:Shot/IsOutdoor"
	ShotIsNight      trimmer.MetaPath = "trim:Shot/IsNight"
	ShotIsDay        trimmer.MetaPath = "trim:Shot/IsDay"
	ShotHasSound     trimmer.MetaPath = "trim:Shot/HasSound"
	ShotTimecodeIn   trimmer.MetaPath = "trim:Shot/TimecodeIn"
	ShotTimecodeOut  trimmer.MetaPath = "trim:Shot/TimecodeOut"
	ShotTimecodeBits trimmer.MetaPath = "trim:Shot/TimecodeBits"
	ShotIsDropFrame  trimmer.MetaPath = "trim:Shot/IsDropFrame"
	ShotDistanceUnit trimmer.MetaPath = "trim:Shot/DistanceUnit"

	SoundId           trimmer.MetaPath = "trim:Sound/Id"
	SoundIndex        trimmer.MetaPath = "trim:Sound/Index"
	SoundMake         trimmer.MetaPath = "trim:Sound/Make"
	SoundModel        trimmer.MetaPath = "trim:Sound/Model"
	SoundSoftware     trimmer.MetaPath = "trim:Sound/Software"
	SoundSerial       trimmer.MetaPath = "trim:Sound/Serial"
	SoundMakeDate     trimmer.MetaPath = "trim:Sound/MakeDate"
	SoundNotes        trimmer.MetaPath = "trim:Sound/Notes"
	SoundRecorderType trimmer.MetaPath = "trim:Sound/RecorderType"
	SoundMicType      trimmer.MetaPath = "trim:Sound/MicType"
	SoundSource       trimmer.MetaPath = "trim:Sound/Source"

	StorageId          trimmer.MetaPath = "trim:Storage/Id"
	StorageIndex       trimmer.MetaPath = "trim:Storage/Index"
	StorageMake        trimmer.MetaPath = "trim:Storage/Make"
	StorageModel       trimmer.MetaPath = "trim:Storage/Model"
	StorageSoftware    trimmer.MetaPath = "trim:Storage/Software"
	StorageSerial      trimmer.MetaPath = "trim:Storage/Serial"
	StorageMakeDate    trimmer.MetaPath = "trim:Storage/MakeDate"
	StorageNotes       trimmer.MetaPath = "trim:Storage/Notes"
	StorageType        trimmer.MetaPath = "trim:Storage/Type"
	StorageCapacity    trimmer.MetaPath = "trim:Storage/Capacity"
	StorageFilesystem  trimmer.MetaPath = "trim:Storage/Filesystem"
	StorageTemperature trimmer.MetaPath = "trim:Storage/Temperature"
)
